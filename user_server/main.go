package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	user "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/connection"
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
	postClient := user.NewPostServiceClient(postServiceConnection)
	// create a new PostServiceClient
	newPost := &user.Post{
		Id:          125,
		Title:       "This is a test title for the user service.",
		Description: "Test Description",
		IsActive:    true,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	createdPost, err := postClient.Create(context.Background(), newPost)
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	log.Printf("Created post: %v", createdPost)

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
