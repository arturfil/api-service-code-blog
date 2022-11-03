package user

import (
	"context"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserStore interface {
	Signup(context.Context, User) (User, error)
	GetUserByEmail(context.Context, string) (User, error)
	PasswordMatches(string, User) (bool, error)
}

type Service struct {
	Store UserStore
}

// struc that the login method would receive as parameter
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewService(store UserStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) Signup(ctx context.Context, user User) (User, error) {
	user, err := s.Store.Signup(ctx, user)
	if err != nil {
		fmt.Println("ERROR", err)
		return User{}, nil
	}
	return user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := s.Store.GetUserByEmail(ctx, email)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	// get user by email
	return user, nil
}

func (s *Service) PasswordMatches(plainText string, user User) (bool, error) {
	isValid, err := s.Store.PasswordMatches(plainText, user)
	if err != nil {
		log.Print(err)
		return false, nil
	}
	return isValid, nil
}
