package book

import (
	"context"
	"errors"
)

var (
	ErrNotFound  = errors.New("book not found")
	ErrDuplicate = errors.New("book already exists")
)

type Repository interface {
	FindByID(ctx context.Context, id string) (Book, error)
	ISBNUsedByOther(ctx context.Context, isbn, id string) (bool, error)
	Update(ctx context.Context, book Book) (Book, error)
	ExistsByISBN(ctx context.Context, isbn string) (bool, error)
	Insert(ctx context.Context, book Book) (Book, error)
}
