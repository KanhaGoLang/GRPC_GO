package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/KanhaGoLang/go_common/common"
	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.GreenString("USER-SERVICE create User %v", user))

	common.MyLogger.Println(color.YellowString("USER-SERVICE checking if User Email already exists %s", user.Email))
	_, err := s.getUserIDByEmail(user.Email)
	if err == nil {
		common.MyLogger.Println(color.RedString("USER-SERVICE User already exists with Email %s", user.Email))

		return nil, errors.New("email already exists")
	}

	query := "INSERT INTO users (name, email, password, role, is_active) VALUES (?, ?, ?, ?, ?)"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result, err := s.db.ExecContext(ctx, query, user.Name, user.Email, hashedPassword, user.Role, user.IsActive)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int32(userId)

	return user, nil
}

func (s *UserService) getUserIDByEmail(email string) (int32, error) {
	query := "SELECT id FROM users where email = ?"

	var user proto.User
	row := s.db.QueryRow(query, email)
	err := row.Scan(&user.Id)
	if err != nil {
		return 0, errors.New("user not found")
	}
	return user.Id, nil

}

func (s *UserService) ReadUser(ctx context.Context, req *proto.UserId) (*proto.User, error) {
	if req.Id > 0 {
		common.MyLogger.Println(color.GreenString("USER-SERVICE get user by Id %v", req.Id))
	}

	if req == nil || req.Id < 0 {
		return nil, fmt.Errorf("invalid id %v", req.Id)
	}

	query := "SELECT * FROM users WHERE id = ?"

	row := s.db.QueryRow(query, req.Id)

	var user proto.User

	// scan the row in the user struct

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		common.MyLogger.Println(color.RedString("error reading from database : %s", err))
		return nil, err
	}

	return &user, nil

}

func (us *UserService) UpdateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.GreenString("USER-SERVICE update User %v", req))
	common.MyLogger.Println(color.YellowString("USER-SERVICE checking if User Email already exists %s", req.Email))

	existingUserIDWIthEmail, err := us.getUserIDByEmail(req.Email)

	if err == nil && existingUserIDWIthEmail != req.Id {
		common.MyLogger.Println(color.RedString("USER-SERVICE User already exists email %s", req.Email))

		return nil, errors.New("email is already used by some other users")
	}

	if req == nil || req.Id <= 0 {
		return nil, fmt.Errorf("invalid payload")
	}

	// creating a type dynamically
	var userId proto.UserId
	userId.Id = req.Id

	// check if user exists in DB
	_, err = us.ReadUser(ctx, &userId)

	if err != nil {
		return nil, err
	}

	req.UpdatedAt = time.Now().Format(time.DateTime)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	query := "UPDATE users SET name = ?, email = ? , password = ?, updated_at = ? WHERE id = ?"

	_, err = us.db.ExecContext(ctx, query, req.Name, req.Email, hashedPassword, req.UpdatedAt, req.Id)

	if err != nil {
		return nil, err
	}
	req.Password = ""

	return req, nil
}

func (us *UserService) GetAllUsers(ctx context.Context, req *proto.NoParameter) (*proto.Users, error) {
	common.MyLogger.Println(color.GreenString("USER-SERVICE get all Users"))

	rows, err := us.db.QueryContext(ctx, "SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	// var user user.User
	users := []*proto.User{}

	for rows.Next() {
		u := new(proto.User)
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &proto.Users{User: users}, nil

}

func (us *UserService) DeleteUser(ctx context.Context, req *proto.UserId) (*proto.UserSuccess, error) {
	common.MyLogger.Println(color.GreenString("USER-SERVICE delete user by Id %d", req.Id))

	query := "DELETE FROM users WHERE id = ?"

	_, err := us.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.UserSuccess{IsSuccess: true}, nil
}
