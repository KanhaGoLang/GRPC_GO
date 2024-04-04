package service

import (
	"context"
	"errors"

	"github.com/KanhaGoLang/grpc_go/proto"
	"google.golang.org/grpc"
)

// UserService defines methods for interacting with the gRPC user service
type UserService interface {
	GetUsers() (*proto.Users, error)
	// Define other methods as needed
}

type userGrpcServiceClient struct {
	grpcClient proto.UserServiceClient
}

// NewUserServiceClient initializes a new instance of UserServiceClient
func NewUserServiceClient(grpcConn *grpc.ClientConn) UserService {
	return &userGrpcServiceClient{grpcClient: proto.NewUserServiceClient(grpcConn)}
}

func (uc *userGrpcServiceClient) GetUsers() (*proto.Users, error) {
	users, err := uc.grpcClient.GetAllUsers(context.Background(), &proto.NoParameter{})
	if err != nil {
		return nil, errors.New("unable to get user")
	}
	return users, nil

}
