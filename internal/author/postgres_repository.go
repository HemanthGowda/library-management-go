package author

import (
	"context"
	"github.com/HemanthGowda/library-management-go/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) ExistByName(ctx context.Context, name string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(
		ctx,
		`SELECT EXISTS(
            SELECT 1 FROM authors
            WHERE name = $1
        )`,
		name,
	).Scan(&exists)

	return exists, err
}

func (r *PostgresRepository) Create(ctx context.Context, author Author) (Author, error) {

	author.Id = uuid.NewString()

	query := `
		INSERT INTO authors (
			id,
			name,
			biography,
			birth,
			nationality
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := r.pool.QueryRow(
		ctx,
		query,
		author.Id,
		author.Name,
		author.Biography,
		author.Birth,
		author.Nationality,
	).Scan(
		&author.Id,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if db.IsUniqueViolation(err) {
		return Author{}, ErrDuplicate
	}

	if err != nil {
		return Author{}, err
	}

	return author, nil
}

func (r *PostgresRepository) ViewAuthor(ctx context.Context, id string) (Author, error) {
	var author Author
	query := `
        SELECT id, name, biography, birth, nationality,
               created_at, updated_at
        FROM authors
        WHERE id = $1
    `
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&author.Id,
		&author.Name,
		&author.Biography,
		&author.Birth,
		&author.Nationality,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		return Author{}, err
	}

	return author, nil
}
