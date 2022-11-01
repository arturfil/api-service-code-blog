package user

import (
	"context"
	"fmt"
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
}

type Service struct {
	Store UserStore
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
