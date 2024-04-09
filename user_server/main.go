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

func main() {
	common.MyLogger.Println(color.CyanString("UserServer is starting..."))
	common.MyLogger.Println(color.HiMagentaString("Connecting to PostGRPC Service on port %s", common.PostServiceAddress))

	// Set up a connection to the PostService server
	postServiceConnection, err := grpc.Dial(common.PostServiceAddress, grpc.WithInsecure())
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

	listener, tcpErr := net.Listen("tcp", common.UserServiceAddress)
	if tcpErr != nil {
		panic(tcpErr)
	}

	// adding grpc option to validate JWT token aby adding an Auth Interceptor
	grpcOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(service.AuthInterceptor),
	}

	grpcServer := grpc.NewServer(grpcOpts...)
	proto.RegisterUserServiceServer(grpcServer, &controller.UserController{UserService: userService, PostServiceClient: postClient})

	common.MyLogger.Println(color.BlueString("UserServer started on port %s", common.UserServiceAddress))

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}

}
