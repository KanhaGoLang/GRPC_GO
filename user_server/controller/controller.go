package controller

import (
	"context"
	"fmt"
	"log"

	"io"

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

func (uc *UserController) SaveMultipleUsers(stream proto.UserService_SaveMultipleUsersServer) error {

	var users []*proto.User
	for {
		user, err := stream.Recv()
		if err != nil {
			// If the stream has ended
			if err == io.EOF {
				// Save all users
				// For demonstration, you can print received users
				fmt.Println(users)
				// Return success response
				return stream.SendAndClose(&proto.UserSuccess{IsSuccess: true})
				// return stream.SendAndClose(&proto.UserSuccess{IsSuccess: true})
			}
			return err
		}
		// Append user to the slice
		users = append(users, user)
		log.Println(users)
	}
}
