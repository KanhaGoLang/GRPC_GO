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
	return nil, fmt.Errorf("not implemented sanjeev")
}
