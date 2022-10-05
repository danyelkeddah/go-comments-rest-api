package main

import (
	"context"
	"fmt"
	"github.com/danyelkeddah/go-comments-rest-api/internal/comment"
	"github.com/danyelkeddah/go-comments-rest-api/internal/database"
)

// Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("starting up our application")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDb(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	commentService := comment.NewService(db)
	fmt.Println(commentService.GetComment(

		context.Background(),
		"e53dcd6b-8c72-4ef7-b9f5-fc9c686ec9b4",
	))

	return nil
}

func main() {
	fmt.Println("Go rest api ")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
