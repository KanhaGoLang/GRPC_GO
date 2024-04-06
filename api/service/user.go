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
	GetUsers() (*proto.Users, error)
	GetUser(id int32) (*proto.User, error)
}

type userGrpcServiceClient struct {
	grpcClient proto.UserServiceClient
}

// NewUserServiceClient initializes a new instance of UserServiceClient
func NewUserServiceClient(grpcConn *grpc.ClientConn) UserService {
	return &userGrpcServiceClient{grpcClient: proto.NewUserServiceClient(grpcConn)}
}

func (us *userGrpcServiceClient) GetUsers() (*proto.Users, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE get all Users"))

	users, err := us.grpcClient.GetAllUsers(context.Background(), &proto.NoParameter{})
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (us *userGrpcServiceClient) GetUser(id int32) (*proto.User, error) {
	common.MyLogger.Println(color.MagentaString("USER-SERVICE get user by userID %d", id))

	user, err := us.grpcClient.ReadUserById(context.Background(), &proto.UserId{Id: id})
	if err != nil {
		return nil, err
	}
	return user, nil
}
