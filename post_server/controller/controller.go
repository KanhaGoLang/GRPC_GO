package controller

import (
	"context"
	"log"

	"github.com/KanhaGoLang/grpc_go/post_server/service"
	proto "github.com/KanhaGoLang/grpc_go/proto"
)

type PostController struct {
	PostService *service.PostService
	proto.UnimplementedPostServiceServer
}

func (pc *PostController) GetAll(ctx context.Context, req *proto.NoPostParameter) (*proto.Posts, error) {
	log.Println("PC get all posts")

	return pc.PostService.GetAll(ctx, req)
}

func (pc *PostController) ReadById(ctx context.Context, req *proto.PostId) (*proto.Post, error) {
	return pc.PostService.ReadById(ctx, req)
}

func (pc *PostController) Delete(ctx context.Context, req *proto.PostId) (*proto.PostSuccess, error) {
	return pc.PostService.Delete(ctx, req)

}

func (pc *PostController) Create(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	return pc.PostService.Create(ctx, req)
}

func (pc *PostController) Update(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	return pc.PostService.Update(ctx, req)
}
