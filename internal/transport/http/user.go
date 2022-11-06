package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arturfil/go_code_blog_api/internal/user"
	"github.com/golang-jwt/jwt/v4"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string    `json:"token"`
	User  user.User `json:"user"`
}

type UserService interface {
	Signup(ctx context.Context, user user.User) (user.User, error)
	// Login(ctx context.Context, email string) (user.User, error)
	GetUserByEmail(ctx context.Context, email string) (user.User, error)
	PasswordMatches(plainText string, user user.User) (bool, error)
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var usr user.User
	var myKey = []byte(os.Getenv("SECRET_KEY"))
	if err := json.NewDecoder(r.Body).Decode(&usr); err != nil {
		return
	}
	// check if users exists
	userToLogIn, err := h.UserService.GetUserByEmail(r.Context(), usr.Email)
	if err != nil {
		log.Print("User with that email wasn't found")
		return
	}
	// check if password matches
	isValid, err := h.UserService.PasswordMatches(usr.Password, userToLogIn)
	if err != nil || !isValid {
		log.Print("Check credentials")
		return
	}
	// create JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["name"] = usr.FirstName
	claims["email"] = usr.Email
	claims["exp"] = time.Now().Add(time.Minute * 60 * 4).Unix() // 4 hours
	// sign token
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		log.Print(err)
		return
	}
	// make sure we are not returning the hashed password
	userToLogIn.Password = "hidden"
	response := tokenResponse{
		Token: tokenString,
		User:  userToLogIn,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
