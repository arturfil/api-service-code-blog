package db

import (
	"context"
	"errors"
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
	Clearance string    `json:"clearance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// method to signup a new user
func (d *Database) Signup(ctx context.Context, user user.User) (user.User, error) {
	newId := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return user, err
	}
	query := `
		insert into users (id, email, first_name, last_name, clearance, password, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id;
	`
	err = d.Client.QueryRowContext(ctx, query,
		newId,
		user.Email,
		user.FirstName,
		user.LastName,
		"member",
		hashedPassword,
		time.Now(),
		time.Now(),
	).Scan(&newId)

	if err != nil {
		return user, err
	}
	return user, nil
}

// method to get user by email
func (d *Database) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	query := `
		select id, email, first_name, last_name, password, created_at, updated_at from users
		where email = $1
	`
	row := d.Client.QueryRowContext(ctx, query, email)
	var user user.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d *Database) PasswordMatches(plainText string, user user.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, err
		default:
			return false, err
		}
	}
	return true, nil
}
