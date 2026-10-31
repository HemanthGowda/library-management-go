package user

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

func (r *PostgresRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM users WHERE lower(email) = lower($1))`, email).Scan(&exists)
	return exists, err
}

func (r *PostgresRepository) Create(ctx context.Context, user User) (User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return User{}, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	user.ID = uuid.NewString()
	err = tx.QueryRow(ctx, `
		INSERT INTO users (id, email, password_hash, full_name, enabled)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at, updated_at
	`, user.ID, user.Email, user.PasswordHash, user.FullName, user.Enabled).Scan(&user.CreatedAt, &user.UpdatedAt)
	if db.IsUniqueViolation(err) {
		return User{}, ErrDuplicate
	}
	if err != nil {
		return User{}, err
	}

	for _, role := range user.Roles {
		if _, err := tx.Exec(ctx, `INSERT INTO user_roles (user_id, role) VALUES ($1, $2)`, user.ID, string(role)); err != nil {
			return User{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	var roles []string
	err := r.pool.QueryRow(ctx, `
		SELECT u.id::text, u.email, u.password_hash, u.full_name, u.enabled, u.created_at, u.updated_at,
		       COALESCE(ARRAY_AGG(ur.role ORDER BY ur.role) FILTER (WHERE ur.role IS NOT NULL), '{}')
		FROM users u
		LEFT JOIN user_roles ur ON ur.user_id = u.id
		WHERE lower(u.email) = lower($1)
		GROUP BY u.id
	`, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FullName,
		&user.Enabled,
		&user.CreatedAt,
		&user.UpdatedAt,
		&roles,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return User{}, ErrNotFound
	}
	if err != nil {
		return User{}, err
	}

	user.Roles = make([]Role, 0, len(roles))
	for _, role := range roles {
		user.Roles = append(user.Roles, Role(role))
	}
	return user, nil
}
