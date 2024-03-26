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
