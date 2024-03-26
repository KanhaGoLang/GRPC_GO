package service

import (
	"context"
	"database/sql"
	"fmt"

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
