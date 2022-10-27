package post

import (
	"context"
	"fmt"
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// here we pass the types of data
type PostStore interface {
	GetPosts(context.Context) ([]Post, error)
	GetPostById(context.Context, string) (Post, error)
	CreatePost(context.Context, Post) (Post, error)
	UpdatePost(context.Context, Post, string) (Post, error)
}

type Service struct {
	Store PostStore
}

func NewService(store PostStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetPosts(ctx context.Context) ([]Post, error) {
	posts, err := s.Store.GetPosts(ctx)
	if err != nil {
		fmt.Println(err)
		return []Post{}, err
	}
	return posts, nil
}

func (s *Service) GetPostById(ctx context.Context, id string) (Post, error) {
	post, err := s.Store.GetPostById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Post{}, err
	}
	return post, nil
}

func (s *Service) CreatePost(ctx context.Context, post Post) (Post, error) {
	post, err := s.Store.CreatePost(ctx, post)
	if err != nil {
		return Post{}, nil
	}
	return post, nil
}

func (s *Service) UpdatePost(ctx context.Context, post Post, id string) (Post, error) {
	post, err := s.Store.UpdatePost(ctx, post, id)
	if err != nil {
		return Post{}, nil
	}
	return post, nil
}
