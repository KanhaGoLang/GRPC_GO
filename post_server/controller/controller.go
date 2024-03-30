package controller

import (
	"context"
	"fmt"

	post "github.com/KanhaGoLang/grpc_go/proto"
)

type PostController struct {
	post.UnimplementedPostServiceServer
}

func (pc *PostController) Create(ctx context.Context, req *post.Post) (*post.Post, error) {
	fmt.Println("PC create func")
	fmt.Printf("%v\n", req)

	return &post.Post{Id: 1, Title: "Test", Description: "This is a test post", IsActive: true}, nil

	// return nil, fmt.Errorf("not implemented sanjeev")
}
