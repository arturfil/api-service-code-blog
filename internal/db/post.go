package db

import (
	"context"
	"log"
	"time"

	"github.com/arturfil/go_code_blog_api/internal/post"
	"github.com/google/uuid"
)

type PostRow struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertPostRowToField(p PostRow) post.Post {
	return post.Post{
		ID:        p.ID,
		Title:     p.Title,
		Author:    p.Author,
		Content:   p.Content,
		Category:  p.Category,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (d *Database) GetPosts(ctx context.Context) ([]post.Post, error) {
	query := `select * from posts`
	rows, err := d.Client.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	var posts []post.Post
	for rows.Next() {
		var post post.Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Author,
			&post.Content,
			&post.Category,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (d *Database) GetPostById(ctx context.Context, id string) (post.Post, error) {
	query := `select * from posts where id = $1`
	row := d.Client.QueryRowContext(ctx, query, id)
	var post post.Post
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Author,
		&post.Content,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (d *Database) CreatePost(ctx context.Context, post post.Post) (post.Post, error) {
	newId := uuid.New()
	query := `
		insert into posts (id, title, author, content, category, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id;
	`

	err := d.Client.QueryRowContext(ctx, query,
		newId,
		post.Title,
		post.Author,
		post.Content,
		post.Category,
		time.Now(),
		time.Now(),
	).Scan(&newId)

	if err != nil {
		log.Print(err)
		return post, nil
	}
	return post, nil
}

func (d *Database) UpdatePost(ctx context.Context, post post.Post, id string) (post.Post, error) {
	query := `
		update posts set 
		title = $1, 
		author = $2, 
		content = $3, 
		category = $4
		where id = $5
	`

	err := d.Client.QueryRowContext(ctx, query,
		post.Title,
		post.Author,
		post.Content,
		post.Category,
		id,
	)
	if err != nil {
		log.Print(err)
		return post, nil
	}
	return post, nil
}
