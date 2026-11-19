package author

import (
	"context"
	"errors"
)

var (
	ErrNotFound  = errors.New("Author not found")
	ErrDuplicate = errors.New("Author already exists")
)

type Repository interface {
	ExistByName(ctx context.Context, name string) (bool, error)
	Create(ctx context.Context, author Author) (Author, error)
	ViewAuthor(ctx context.Context, id string) (Author, error)
	SearchAuthor(ctx context.Context, name string, page int, limit int) ([]Author, error)
	CountAuthors(ctx context.Context, name string) (int, error)
}
