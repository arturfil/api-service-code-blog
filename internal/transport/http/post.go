package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arturfil/go_code_blog_api/internal/post"
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

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post post.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return
	}
	post, err := h.PostService.CreatePost(r.Context(), post)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(post); err != nil {
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
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}
