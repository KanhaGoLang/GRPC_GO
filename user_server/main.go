package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	user "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/connection"
	"github.com/KanhaGoLang/grpc_go/user_server/controller"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Initialize database connection
	db, err := connection.NewDatabaseConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("Connected to Database")
	}
	defer db.Close()

	// Initialize UserService
	userService := service.NewUserService(db)

	listener, tcpErr := net.Listen("tcp", "localhost:50052")
	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, &controller.UserController{UserService: userService})

	fmt.Println("Server started")

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

}
