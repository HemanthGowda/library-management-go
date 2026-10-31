package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/HemanthGowda/library-management-go/internal/auth"
	"github.com/HemanthGowda/library-management-go/internal/book"
	"github.com/HemanthGowda/library-management-go/internal/domain"
	"github.com/HemanthGowda/library-management-go/internal/httpx"
	"github.com/HemanthGowda/library-management-go/internal/user"
	"github.com/HemanthGowda/library-management-go/openapi"
)

type Dependencies struct {
	Ping   func(context.Context) error
	Tokens *auth.JWTService
	Users  *user.Handler
	Auth   *auth.Handler
	Books  *book.Handler
}

func NewRouter(deps Dependencies) (http.Handler, error) {
	spec, err := openapi.JSON()
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()
	router.Use(recoverJSON)
	router.Use(auth.Middleware(deps.Tokens))

	router.Get("/actuator/health", healthHandler(deps.Ping))
	router.Get("/actuator/info", infoHandler)
	router.Get("/v3/api-docs", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(spec)
	})
	router.Get("/swagger-ui.html", swaggerUI)
	router.Post("/api/v1/auth/login", httpx.Handle(deps.Auth.Login))
	router.Post("/api/v1/users", httpx.Handle(deps.Users.Register))
	router.Put("/api/v1/books/{id}", httpx.Handle(deps.Books.Update))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := auth.PrincipalFrom(r.Context()); !ok {
			httpx.WriteError(w, r, domain.Unauthorized("Authentication is required"))
			return
		}
		httpx.WriteError(w, r, domain.NotFound("Resource not found"))
	})

	return router, nil
}

func healthHandler(ping func(context.Context) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		status := http.StatusOK
		body := map[string]string{"status": "UP"}
		if err := ping(ctx); err != nil {
			status = http.StatusServiceUnavailable
			body["status"] = "DOWN"
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(body)
	}
}

func infoHandler(w http.ResponseWriter, _ *http.Request) {
	_ = httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"name":    "library-management-go",
		"version": "0.1.0",
	})
}

func swaggerUI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8"/>
  <title>Library Management API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui.css"/>
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@5.17.14/swagger-ui-bundle.js"></script>
<script>
window.ui = SwaggerUIBundle({
  url: "/v3/api-docs",
  dom_id: "#swagger-ui",
  persistAuthorization: true
});
</script>
</body>
</html>`))
}

func recoverJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovered := recover(); recovered != nil {
				slog.Error("panic", "error", recovered, "stack", string(debug.Stack()))
				httpx.WriteError(w, r, errors.New("panic"))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
