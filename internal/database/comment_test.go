//go:build integration
// +build integration

package database

import (
	"context"
	"fmt"
	"github.com/danyelkeddah/go-comments-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Body:   "body",
			Author: "author",
		})

		assert.NoError(t, err)

		newComment, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newComment.Slug)

		fmt.Println("testing the creation of comment")
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Body:   "new-body",
			Author: "new-author",
		})

		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)

		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)

	})

	t.Run("test get comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Body:   "new-body",
			Author: "new-author",
		})
		assert.NoError(t, err)

		comment, err := db.GetComment(context.Background(), cmt.ID)

		assert.NoError(t, err)
		assert.Equal(t, "new-slug", comment.Slug)

	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Body:   "new-body",
			Author: "new-author",
		})
		assert.NoError(t, err)
		updatedComment, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "updated-slug",
			Body:   "updated-body",
			Author: "updated-author",
		})
		assert.NoError(t, err)
		assert.Equal(t, "updated-slug", updatedComment.Slug)

	})
}
