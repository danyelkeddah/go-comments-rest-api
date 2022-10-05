package comment

import "context"

// Repository - (Store) This interface defines all the methods that our service need in order to operate
type Repository interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
	Ping(ctx context.Context) error
}
