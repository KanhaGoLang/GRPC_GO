package controller

import (
	"context"
	"fmt"

	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
)

type UserController struct {
	UserService *service.UserService
	proto.UnimplementedUserServiceServer
	PostServiceClient proto.PostServiceClient
}

func (uc *UserController) ReadUserById(ctx context.Context, req *proto.UserId) (*proto.User, error) {
	fmt.Println("Server ReadUserById")

	dbUser, err := uc.UserService.ReadUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return dbUser, nil
}

func (uc *UserController) CreateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	fmt.Println("Server create user")

	createdUser, err := uc.UserService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(createdUser)

	return createdUser, nil
}

func (uc *UserController) UpdateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	fmt.Println("UC update user")

	return uc.UserService.UpdateUser(ctx, req)
}

func (uc *UserController) GetAllUsers(ctx context.Context, req *proto.NoParameter) (*proto.Users, error) {
	fmt.Println("UC get all users")

	return uc.UserService.GetAllUsers(ctx, req)
}

func (uc *UserController) DeleteUser(ctx context.Context, req *proto.UserId) (*proto.UserSuccess, error) {
	fmt.Println("UC Delete user")

	return uc.UserService.DeleteUser(ctx, req)
}
