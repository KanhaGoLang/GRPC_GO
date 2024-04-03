package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KanhaGoLang/go_common/common"
	"google.golang.org/grpc"

	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/controller"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
	_ "github.com/go-sql-driver/mysql"
)

const (
	postServiceAddress = "localhost:50053" // Assuming PostService runs on localhost:50051
)

func main() {
	// Set up a connection to the PostService server
	postServiceConnection, err := grpc.Dial(postServiceAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer postServiceConnection.Close()

	// Create a PostService client
	postClient := proto.NewPostServiceClient(postServiceConnection)

	// Initialize database connection
	db, err := common.NewDatabaseConnection()
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
	proto.RegisterUserServiceServer(grpcServer, &controller.UserController{UserService: userService, PostServiceClient: postClient})

	fmt.Println("Server started")

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

}
