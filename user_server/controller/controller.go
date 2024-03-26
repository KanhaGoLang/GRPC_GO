package controller

import (
	"context"
	"fmt"

	user "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
)

type UserController struct {
	UserService *service.UserService
	user.UnimplementedUserServiceServer
}

func (u *UserController) ReadUserById(ctx context.Context, req *user.UserId) (*user.User, error) {
	fmt.Println("Server ReadUserById")

	dbUser, err := u.UserService.ReadUser(ctx, req)

	if err != nil {
		return nil, err
	}

	fmt.Println(dbUser)

	return dbUser, nil

	// return &user.User{
	// 	Id:       1,
	// 	Name:     "Sanjeev",
	// 	Email:    "test@test.com",
	// 	Password: "test",
	// 	Role:     "admin",
	// 	IsActive: true,
	// }, nil
}

func (u *UserController) CreateUser(ctx context.Context, req *user.User) (*user.User, error) {
	fmt.Println("Server create user")

	createdUser, err := u.UserService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(createdUser)

	return createdUser, nil
}
