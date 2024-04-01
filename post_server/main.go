package main

import (
	"fmt"
	"net"

	"github.com/KanhaGoLang/grpc_go/common"
	"github.com/KanhaGoLang/grpc_go/post_server/controller"
	"github.com/KanhaGoLang/grpc_go/post_server/service"
	post "github.com/KanhaGoLang/grpc_go/proto"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Initialize database connection
	db, err := common.NewDatabaseConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("Connected to Database")
	}
	defer db.Close()

	listener, tcpErr := net.Listen("tcp", "localhost:50053")

	if tcpErr != nil {
		panic(tcpErr)
	}

	//initialize postService
	postService := service.NewPostService(db)

	grpcServer := grpc.NewServer()

	post.RegisterPostServiceServer(grpcServer, &controller.PostController{PostService: postService})

	fmt.Println("Starting server")

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}
}
