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

func (uc *UserController) ReadUserById(ctx context.Context, req *user.UserId) (*user.User, error) {
	fmt.Println("Server ReadUserById")

	dbUser, err := uc.UserService.ReadUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return dbUser, nil
}

func (uc *UserController) CreateUser(ctx context.Context, req *user.User) (*user.User, error) {
	fmt.Println("Server create user")

	createdUser, err := uc.UserService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(createdUser)

	return createdUser, nil
}

func (uc *UserController) UpdateUser(ctx context.Context, req *user.User) (*user.User, error) {
	fmt.Println("UC update user")

	return uc.UserService.UpdateUser(ctx, req)
}

func (uc *UserController) GetAllUsers(ctx context.Context, req *user.NoParameter) (*user.Users, error) {
	fmt.Println("UC get all users")

	return uc.UserService.GetAllUsers(ctx, req)
	// return nil, errors.New("function not implemented")
}
