package service

import (
	"context"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
	"google.golang.org/grpc"
)

// UserService defines methods for interacting with the gRPC user service
type UserService interface {
	GetUsers(ctx context.Context) (*proto.Users, error)
	GetUser(ctx context.Context, id int32) (*proto.User, error)
	CreateUser(ctx context.Context, user *proto.User) (*proto.User, error)
	UpdateUser(ctx context.Context, user *proto.User) (*proto.User, error)
	Delete(ctx context.Context, id int32) (*proto.UserSuccess, error)
}

type userGrpcServiceClient struct {
	grpcClient proto.UserServiceClient
}

// NewUserServiceClient initializes a new instance of UserServiceClient
func NewUserServiceClient(grpcConn *grpc.ClientConn) UserService {
	return &userGrpcServiceClient{grpcClient: proto.NewUserServiceClient(grpcConn)}
}

func (us *userGrpcServiceClient) GetUsers(ctx context.Context) (*proto.Users, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE get all Users"))

	users, err := us.grpcClient.GetAllUsers(ctx, &proto.NoParameter{})
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (us *userGrpcServiceClient) GetUser(ctx context.Context, id int32) (*proto.User, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE get user by userID %d", id))

	user, err := us.grpcClient.ReadUserById(ctx, &proto.UserId{Id: id})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userGrpcServiceClient) CreateUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE create User %v", user))

	newUser, err := us.grpcClient.CreateUser(ctx, &proto.User{Name: user.Name, Email: user.Email, Password: user.Password, Role: user.Role, IsActive: user.IsActive})

	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (us *userGrpcServiceClient) UpdateUser(ctx context.Context, user *proto.User) (*proto.User, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE update User %v", user))

	updatedUser, err := us.grpcClient.UpdateUser(ctx, &proto.User{Id: user.Id, Name: user.Name, Email: user.Email, Password: user.Password, Role: user.Role, IsActive: user.IsActive})

	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (us *userGrpcServiceClient) Delete(ctx context.Context, id int32) (*proto.UserSuccess, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE delete User %v", id))

	updatedUser, err := us.grpcClient.DeleteUser(ctx, &proto.UserId{Id: id})

	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
