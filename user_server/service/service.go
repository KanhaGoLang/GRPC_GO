package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	user "github.com/KanhaGoLang/grpc_go/proto"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	fmt.Println("Service create user")

	query := "INSERT INTO users (name, email, password, role, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := s.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Role, user.IsActive, user.CreatedAt, user.UpdatedAt)
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

func (s *UserService) ReadUser(ctx context.Context, req *user.UserId) (*user.User, error) {
	if req == nil || req.Id < 0 {
		return nil, fmt.Errorf("invalid id %v", req.Id)
	}

	query := "SELECT * FROM users WHERE id = ?"

	row := s.db.QueryRow(query, req.Id)

	var user user.User

	// scan the row in the user struct

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (us *UserService) UpdateUser(ctx context.Context, req *user.User) (*user.User, error) {
	fmt.Println("US Update user")

	if req == nil || req.Id <= 0 {
		return nil, fmt.Errorf("invalid payload")
	}

	// creating a type dynamically
	var userId user.UserId
	userId.Id = req.Id

	// check if user exists in DB
	_, err := us.ReadUser(ctx, &userId)

	if err != nil {
		return nil, err
	}

	req.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	query := "UPDATE users SET name = ?, email = ? , password = ?, updated_at = ? WHERE id = ?"

	_, err = us.db.ExecContext(ctx, query, req.Name, req.Email, req.Password, req.UpdatedAt, req.Id)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (us *UserService) GetAllUsers(ctx context.Context, req *user.NoParameter) (*user.Users, error) {
	rows, err := us.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	// var user user.User
	users := []*user.User{}

	for rows.Next() {
		u := new(user.User)
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Role, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return &user.Users{User: users}, nil

}

func (us *UserService) DeleteUser(ctx context.Context, req *user.UserId) (*user.UserSuccess, error) {
	query := "DELETE FROM users WHERE id = ?"

	_, err := us.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &user.UserSuccess{IsSuccess: true}, nil
}
