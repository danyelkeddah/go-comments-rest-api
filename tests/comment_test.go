//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("thisisademotestkey"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

func TestCreateComment(t *testing.T) {
	t.Run("can create comment", func(t *testing.T) {
		client := resty.New()
		token := createToken()
		resp, err := client.R().
			SetHeader("Authorization", "Bearer "+token).
			SetBody(`{"slug": "/", "author": "user", "body":"hello world"}`).
			Post("http://localhost:8080/api/v1/comments")

		assert.NoError(t, err)

		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("cannot create comment without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().SetBody(`{"slug": "/", "author": "user", "body":"hello world"}`).Post("http://localhost:8080/api/v1/comments")
		assert.NoError(t, err)

		assert.Equal(t, 401, resp.StatusCode())
	})
}
