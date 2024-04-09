package controller

import (
	"context"
	"log"

	"io"

	"github.com/KanhaGoLang/go_common/common"
	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
	"github.com/fatih/color"
)

type UserController struct {
	UserService *service.UserService
	proto.UnimplementedUserServiceServer
	PostServiceClient proto.PostServiceClient
}

func (uc *UserController) ReadUserById(ctx context.Context, req *proto.UserId) (*proto.User, error) {
	common.MyLogger.Println(color.YellowString("UC--GRPC get user by ID"))

	dbUser, err := uc.UserService.ReadUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return dbUser, nil
}

func (uc *UserController) CreateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.YellowString("UC Server create user %v", req))

	createdUser, err := uc.UserService.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	common.MyLogger.Println(color.MagentaString("Uc createdUser %v", createdUser))

	return createdUser, nil
}

func (uc *UserController) UpdateUser(ctx context.Context, req *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.YellowString("UC update user %v", req))

	return uc.UserService.UpdateUser(ctx, req)
}

func (uc *UserController) GetAllUsers(ctx context.Context, req *proto.NoParameter) (*proto.Users, error) {
	common.MyLogger.Println(color.YellowString("UC get all users"))

	return uc.UserService.GetAllUsers(ctx, req)
}

func (uc *UserController) DeleteUser(ctx context.Context, req *proto.UserId) (*proto.UserSuccess, error) {
	common.MyLogger.Println(color.YellowString("UC Delete user %d", req.Id))

	return uc.UserService.DeleteUser(ctx, req)
}

func (uc *UserController) AuthUser(ctx context.Context, req *proto.AuthRequest) (*proto.TokenResponse, error) {
	common.MyLogger.Println(color.YellowString("UC Authenticate user %s", req.Email))

	return uc.UserService.AuthUser(ctx, req)
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
				common.MyLogger.Println(color.HiMagentaString("%v", users))
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
