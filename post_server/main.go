package main

import (
	"fmt"
	"net"

	"github.com/KanhaGoLang/grpc_go/post_server/controller"
	post "github.com/KanhaGoLang/grpc_go/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, tcpErr := net.Listen("tcp", "localhost:50053")

	if tcpErr != nil {
		panic(tcpErr)
	}

	grpcServer := grpc.NewServer()

	post.RegisterPostServiceServer(grpcServer, &controller.PostController{})

	fmt.Println("Starting server")

	if e := grpcServer.Serve(listener); e != nil {
		panic(e)
	}
}
