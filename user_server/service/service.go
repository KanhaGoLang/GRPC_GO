package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func myLogger() *log.Logger {
	logger := log.New(color.Output, "", 0)

	return logger
}

func (s *UserService) CreateUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	myLogger().Println(color.GreenString("USER-SERVICE create User %v", user))

	query := "INSERT INTO users (name, email, password, role, is_active) VALUES (?, ?, ?, ?, ?)"

	result, err := s.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Role, user.IsActive)
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

func (s *UserService) ReadUser(ctx context.Context, req *proto.UserId) (*proto.User, error) {
	myLogger().Println(color.GreenString("USER-SERVICE get user by Id"))

	if req == nil || req.Id < 0 {
		return nil, fmt.Errorf("invalid id %v", req.Id)
	}

	query := "SELECT * FROM users WHERE id = ?"

	row := s.db.QueryRow(query, req.Id)

	var user proto.User

	// scan the row in the user struct

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		myLogger().Println(color.RedString("error reading from database : %s", err))
		return nil, err
	}

	return &user, nil

}

func (us *UserService) UpdateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	fmt.Println("US Update user")

	if req == nil || req.Id <= 0 {
		return nil, fmt.Errorf("invalid payload")
	}

	// creating a type dynamically
	var userId proto.UserId
	userId.Id = req.Id

	// check if user exists in DB
	_, err := us.ReadUser(ctx, &userId)

	if err != nil {
		return nil, err
	}

	req.UpdatedAt = time.Now().Format(time.DateTime)

	query := "UPDATE users SET name = ?, email = ? , password = ?, updated_at = ? WHERE id = ?"

	_, err = us.db.ExecContext(ctx, query, req.Name, req.Email, req.Password, req.UpdatedAt, req.Id)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (us *UserService) GetAllUsers(ctx context.Context, req *proto.NoParameter) (*proto.Users, error) {
	log.Println("USER Service Get All Users")

	rows, err := us.db.QueryContext(ctx, "SELECT * FROM users")
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
	query := "DELETE FROM users WHERE id = ?"

	_, err := us.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.UserSuccess{IsSuccess: true}, nil
}
