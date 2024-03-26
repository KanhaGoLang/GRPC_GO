package main

import (
	"context"
	"fmt"
	"net"

	user "github.com/KanhaGoLang/grpc_go/proto"
	"google.golang.org/grpc"
)

type userServer struct {
	user.UnimplementedUserServiceServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", "localhost:50052")
	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServer{})

	fmt.Println("Server started")

	if e := grpcServer.Serve(listener); e != nil {
		fmt.Println("panic started")

		panic(e)
	}

	fmt.Println("Server started")

}

func (u *userServer) ReadUserById(ctx context.Context, req *user.UserId) (*user.User, error) {
	fmt.Println("Server ReadUserById")

	return &user.User{
		Id:       1,
		Name:     "Sanjeev",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
		IsActive: true,
	}, nil
}
