package service

import (
	"context"
	"database/sql"
	"log"

	proto "github.com/KanhaGoLang/grpc_go/proto"
)

type PostService struct {
	db *sql.DB
}

func NewPostService(db *sql.DB) *PostService {
	return &PostService{db: db}
}

func (ps *PostService) GetAll(ctx context.Context, req *proto.NoPostParameter) (*proto.Posts, error) {
	log.Println("PS get all posts")

	query := "SELECT * FROM posts"
	log.Println("PS get all posts 23")

	rows, err := ps.db.QueryContext(ctx, query)
	if err != nil {
		println(err.Error()) // proper error handling here

		return nil, err
	}
	log.Println("PS get all posts 32")

	// create posts variable
	posts := []*proto.Post{}
	for rows.Next() {

		post := new(proto.Post)
		err = rows.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			println(err.Error()) // proper error handling here
			return nil, err
		}

		posts = append(posts, post)
	}

	return &proto.Posts{Post: posts}, nil
}
