package book

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/HemanthGowda/library-management-go/internal/db"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) FindByID(ctx context.Context, id string) (Book, error) {
	if _, err := uuid.Parse(id); err != nil {
		return Book{}, ErrNotFound
	}

	row := r.pool.QueryRow(ctx, `
		SELECT id::text, title, author_name, isbn, description, publication_year,
		       total_copies, available_copies, created_at, updated_at
		FROM books
		WHERE id = $1
	`, id)
	return scanBook(row)
}

func (r *PostgresRepository) ISBNUsedByOther(ctx context.Context, isbn, id string) (bool, error) {
	var used bool
	err := r.pool.QueryRow(ctx, `
		SELECT EXISTS(SELECT 1 FROM books WHERE isbn = $1 AND id <> $2)
	`, isbn, id).Scan(&used)
	return used, err
}

func (r *PostgresRepository) Update(ctx context.Context, book Book) (Book, error) {
	row := r.pool.QueryRow(ctx, `
		UPDATE books
		SET title = $2,
		    author_name = $3,
		    isbn = $4,
		    description = $5,
		    publication_year = $6,
		    total_copies = $7,
		    available_copies = $8,
		    updated_at = now()
		WHERE id = $1
		RETURNING id::text, title, author_name, isbn, description, publication_year,
		          total_copies, available_copies, created_at, updated_at
	`, book.ID, book.Title, book.AuthorName, book.ISBN, nullable(book.Description),
		book.PublicationYear, book.TotalCopies, book.AvailableCopies)

	updated, err := scanBook(row)
	if errors.Is(err, ErrNotFound) {
		return Book{}, ErrNotFound
	}
	if db.IsUniqueViolation(err) {
		return Book{}, ErrDuplicate
	}
	return updated, err
}

func (r *PostgresRepository) ExistsByISBN(ctx context.Context, isbn string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM books WHERE isbn = $1)`, isbn).Scan(&exists)
	return exists, err
}

func (r *PostgresRepository) Insert(ctx context.Context, book Book) (Book, error) {
	if book.ID == "" {
		book.ID = uuid.NewString()
	}

	row := r.pool.QueryRow(ctx, `
		INSERT INTO books (
			id, title, author_name, isbn, description, publication_year, total_copies, available_copies
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id::text, title, author_name, isbn, description, publication_year,
		          total_copies, available_copies, created_at, updated_at
	`, book.ID, book.Title, book.AuthorName, book.ISBN, nullable(book.Description),
		book.PublicationYear, book.TotalCopies, book.AvailableCopies)

	inserted, err := scanBook(row)
	if db.IsUniqueViolation(err) {
		return Book{}, ErrDuplicate
	}
	return inserted, err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanBook(row scanner) (Book, error) {
	var book Book
	var description *string
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.AuthorName,
		&book.ISBN,
		&description,
		&book.PublicationYear,
		&book.TotalCopies,
		&book.AvailableCopies,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return Book{}, ErrNotFound
	}
	if err != nil {
		return Book{}, err
	}
	if description != nil {
		book.Description = *description
	}
	book.CreatedAt = book.CreatedAt.UTC()
	book.UpdatedAt = book.UpdatedAt.UTC()
	return book, nil
}

func nullable(value string) any {
	if value == "" {
		return nil
	}
	return value
}
