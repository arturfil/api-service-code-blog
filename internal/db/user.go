package db

import (
	"context"
	"time"

	"github.com/arturfil/go_code_blog_api/internal/user"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRow struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Database) Signup(ctx context.Context, user user.User) (user.User, error) {
	newId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return user, err
	}
	query := `
		insert into users (id, email, first_name, last_name, password, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id;
	`
	err = d.Client.QueryRowContext(ctx, query,
		newId,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		time.Now(),
		time.Now(),
	).Scan(&newId)

	if err != nil {
		return user, err
	}
	return user, nil
}
