package comment

import (
	"context"
	"fmt"
)

// Service - is the struct on which all our logic will be built on top of
type Service struct {
	Repository Repository
}

// NewService - returns a pointer to a new service
func NewService(repository Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieve a comment")
	cmt, err := s.Repository.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, updatedComment Comment) (Comment, error) {
	cmt, err := s.Repository.UpdateComment(ctx, id, updatedComment)

	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}

	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Repository.DeleteComment(ctx, id)
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedComment, err := s.Repository.CreateComment(ctx, cmt)

	if err != nil {
		return Comment{}, err
	}
	return insertedComment, nil
}
