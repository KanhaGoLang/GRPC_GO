package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"google.golang.org/grpc"

	user "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/server/connection"
	"github.com/KanhaGoLang/grpc_go/server/service"
	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db}
}

type userServer struct {
	userService *UserService
	user.UnimplementedUserServiceServer
}

func main() {

	// Initialize database connection
	db, err := connection.NewDatabaseConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Initialize UserService
	userService := NewUserService(db)

	listener, tcpErr := net.Listen("tcp", "localhost:50052")
	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &userServer{userService: userService})

	fmt.Println("Server started")

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

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

func (u *userServer) CreateUser(ctx context.Context, req *user.User) (*user.User, error) {
	fmt.Println("Server create user")

	createdUser, err := service.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Println(createdUser)

	return createdUser, nil
}
