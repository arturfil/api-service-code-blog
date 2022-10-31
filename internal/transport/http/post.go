package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/go_code_blog_api/internal/post"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]post.Post, error)
	GetPostById(ctx context.Context, id string) (post.Post, error)
	CreatePost(ctx context.Context, post post.Post) (post.Post, error)
	UpdatePost(ctx context.Context, post post.Post, id string) (post.Post, error)
}

type Message struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	all, err := h.PostService.GetPosts(r.Context())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(all)
	w.Header().Set("Content-Type", "application/json")
	encjson, _ := json.Marshal(all)
	w.Write(encjson)
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	post, err := h.PostService.GetPostById(r.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(post)
	w.Header().Set("Content-Type", "application/json")
	encjson, _ := json.Marshal(post)
	w.Write(encjson)
}

type CreatePostRequest struct {
	Title    string `json:"title" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Category string `json:"category" validate:"required"`
}

func convertPostRequestToPost(p CreatePostRequest) post.Post {
	return post.Post{
		Title:    p.Title,
		Author:   p.Author,
		Content:  p.Content,
		Category: p.Category,
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return
	}
	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		http.Error(w, "not a valid post", http.StatusBadRequest)
		return
	}

	convertedPost := convertPostRequestToPost(post)
	createdPost, err := h.PostService.CreatePost(r.Context(), convertedPost)

	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var post post.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return
	}
	post, err := h.PostService.UpdatePost(r.Context(), post, id)
	if err != nil {
		log.Println("ERROR", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}
