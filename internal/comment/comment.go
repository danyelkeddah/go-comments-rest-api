package comment

import "errors"

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure of our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}
