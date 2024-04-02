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

func (ps *PostService) ReadById(ctx context.Context, req *proto.PostId) (*proto.Post, error) {
	log.Println("PS get POST by ID")
	query := "SELECT * FROM posts where id = ?"
	row := ps.db.QueryRowContext(ctx, query, req.Id)

	post := new(proto.Post)
	err := row.Scan(&post.Id, &post.Title, &post.Description, &post.IsActive, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostService) Delete(ctx context.Context, req *proto.PostId) (*proto.PostSuccess, error) {
	query := "DELETE FROM posts WHERE id = ?"
	_, err := ps.db.ExecContext(ctx, query, req.Id)

	if err != nil {
		return &proto.PostSuccess{IsSuccess: false}, err
	}

	return &proto.PostSuccess{IsSuccess: true}, nil
}

func (ps *PostService) Create(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	log.Println("PS create POST")

	query := "INSERT INTO posts (title, description, is_active, user_id) VALUES (? , ?, ?, ?)"

	result, err := ps.db.ExecContext(ctx, query, req.Title, req.Description, req.IsActive, req.UserId)

	if err != nil {
		return nil, err
	}

	newPosTtId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return ps.ReadById(ctx, &proto.PostId{Id: int32(newPosTtId)})
}

func (ps *PostService) Update(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	log.Println("PS UPDATE POST")

	dbPost, err := ps.ReadById(ctx, &proto.PostId{Id: req.Id})

	if err != nil {
		return nil, err
	}

	query := "UPDATE posts set title = ?, description = ?, is_active = ?, user_id = ?"

	_, err = ps.db.ExecContext(ctx, query, req.Title, req.Description, req.IsActive, req.UserId)
	if err != nil {
		return nil, err
	}

	dbPost.Title = req.Title
	dbPost.Description = req.Description
	dbPost.IsActive = req.IsActive
	dbPost.UserId = req.UserId

	return dbPost, err
}
