package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/HemanthGowda/library-management-go/internal/auth"
	"github.com/HemanthGowda/library-management-go/internal/book"
	"github.com/HemanthGowda/library-management-go/internal/config"
	"github.com/HemanthGowda/library-management-go/internal/db"
	"github.com/HemanthGowda/library-management-go/internal/httpserver"
	"github.com/HemanthGowda/library-management-go/internal/password"
	"github.com/HemanthGowda/library-management-go/internal/seed"
	"github.com/HemanthGowda/library-management-go/internal/user"
)

type Application struct {
	Handler http.Handler
	Pool    *pgxpool.Pool
	Config  config.Config
	users   user.Repository
	books   book.Repository
}

func New(ctx context.Context, cfg config.Config) (*Application, error) {
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	if err := db.Migrate(cfg.DatabaseURL); err != nil {
		pool.Close()
		return nil, err
	}

	tokens, err := auth.NewJWTService(cfg.JWTSecret, cfg.JWTIssuer, cfg.JWTExpiration)
	if err != nil {
		pool.Close()
		return nil, err
	}

	users := user.NewPostgresRepository(pool)
	books := book.NewPostgresRepository(pool)
	encoder := password.Bcrypt{}
	userService := user.NewService(users, encoder)
	authService := auth.NewService(users, encoder, tokens)
	bookService := book.NewService(books)

	handler, err := httpserver.NewRouter(httpserver.Dependencies{
		Ping:   pool.Ping,
		Tokens: tokens,
		Users:  user.NewHandler(userService),
		Auth:   auth.NewHandler(authService),
		Books:  book.NewHandler(bookService),
	})
	if err != nil {
		pool.Close()
		return nil, err
	}

	return &Application{
		Handler: handler,
		Pool:    pool,
		Config:  cfg,
		users:   users,
		books:   books,
	}, nil
}

func (a *Application) Seed(ctx context.Context) error {
	return seed.Run(ctx, a.Config, a.users, a.books, password.Bcrypt{})
}

func (a *Application) Close() {
	if a.Pool != nil {
		a.Pool.Close()
	}
}
