package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/KanhaGoLang/grpc_go/post_server/service"
	proto "github.com/KanhaGoLang/grpc_go/proto"
)

type PostController struct {
	PostService *service.PostService
	proto.UnimplementedPostServiceServer
}

func (pc *PostController) Create(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	fmt.Println("PC create func")
	fmt.Printf("%v\n", req)

	return &proto.Post{Id: 1, Title: "Test", Description: "This is a test post", IsActive: true}, nil
}

func (pc *PostController) GetAll(ctx context.Context, req *proto.NoPostParameter) (*proto.Posts, error) {
	log.Println("PC get all posts")

	return pc.PostService.GetAll(ctx, req)
}
