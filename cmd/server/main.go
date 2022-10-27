package main

import (
	"fmt"

	"github.com/arturfil/go_code_blog_api/internal/db"
	"github.com/arturfil/go_code_blog_api/internal/post"
	transportHttp "github.com/arturfil/go_code_blog_api/internal/transport/http"
)

// Run - is going to be responsible for the
// instantiation and startup of our application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to db")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	postService := post.NewService(db)

	postHandler := transportHttp.NewHandler(postService)

	if err := postHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("API working")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
