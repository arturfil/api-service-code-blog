package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router      *mux.Router
	PostService PostService
	Server      *http.Server
}

func NewHandler(post_service PostService) *Handler {
	h := &Handler{
		PostService: post_service,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Router.Use(CorsMiddleware)
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LogginMiddleware)
	h.Router.Use(TimeOutMiddleware)
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/posts", h.GetPosts).Methods("GET")
	h.Router.HandleFunc("/api/v1/posts/post", JWTAuh(h.CreatePost)).Methods("POST")
	h.Router.HandleFunc("/api/v1/posts/post/{id}", h.GetPostById).Methods("GET")
	h.Router.HandleFunc("/api/v1/posts/post/{id}", h.UpdatePost).Methods("PUT", "OPTIONS")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
	log.Println("Shutdown gracefully")

	return nil
}
