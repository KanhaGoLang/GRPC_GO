package main

import (
	"net"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/KanhaGoLang/grpc_go/post_server/controller"
	"github.com/KanhaGoLang/grpc_go/post_server/service"
	post "github.com/KanhaGoLang/grpc_go/proto"
	"github.com/fatih/color"
	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Initialize database connection
	db, err := common.NewDatabaseConnection()
	if err != nil {
		common.MyLogger.Println(color.RedString("Error connecting to the database: %v", err))
		return
	} else {
		common.MyLogger.Println(color.GreenString("Connected to Database"))
	}
	defer db.Close()

	listener, tcpErr := net.Listen("tcp", common.PostServiceAddress)

	if tcpErr != nil {
		panic(tcpErr)
	}

	//initialize postService
	postService := service.NewPostService(db)

	grpcServer := grpc.NewServer()

	post.RegisterPostServiceServer(grpcServer, &controller.PostController{PostService: postService})

	common.MyLogger.Println(color.GreenString("POST GRPC Service running on %s", common.PostServiceAddress))

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}
}
