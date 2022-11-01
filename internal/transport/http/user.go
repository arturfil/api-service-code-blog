package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arturfil/go_code_blog_api/internal/user"
)

type UserSerice interface {
	Signup(ctx context.Context, user user.User) (user.User, error)
}

type UserService interface {
	Signup(ctx context.Context, user user.User) (user.User, error)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var usr user.User
	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return
	}
	fmt.Println("USER", usr)

	usr, err := h.UserService.Signup(r.Context(), usr)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		panic(err)
	}
}
