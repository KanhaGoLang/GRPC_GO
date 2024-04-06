package main

import (
	"log"
	"net"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/fatih/color"
	"google.golang.org/grpc"

	proto "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/KanhaGoLang/grpc_go/user_server/controller"
	"github.com/KanhaGoLang/grpc_go/user_server/service"
	_ "github.com/go-sql-driver/mysql"
)

const (
	userServiceAddress = "localhost:50052" // Assuming PostService runs on localhost:50051
	postServiceAddress = "localhost:50053" // Assuming PostService runs on localhost:50051
)

func main() {
	common.MyLogger.Println(color.CyanString("UserServer is starting..."))
	common.MyLogger.Println(color.HiMagentaString("Connecting to PostGRPC Service on port %s", postServiceAddress))

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
		common.MyLogger.Println(color.RedString("Error connecting to the database:", err))

		return
	} else {
		common.MyLogger.Println(color.GreenString("Connected to Database"))
	}
	defer db.Close()

	// Initialize UserService
	userService := service.NewUserService(db)

	listener, tcpErr := net.Listen("tcp", userServiceAddress)
	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, &controller.UserController{UserService: userService, PostServiceClient: postClient})

	common.MyLogger.Println(color.BlueString("UserServer started on port %s", userServiceAddress))

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

}
