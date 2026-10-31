package httpx

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/HemanthGowda/library-management-go/internal/domain"
)

type apiError struct {
	Timestamp  time.Time          `json:"timestamp"`
	Status     int                `json:"status"`
	Error      string             `json:"error"`
	Message    string             `json:"message"`
	Path       string             `json:"path"`
	Violations []domain.Violation `json:"violations,omitempty"`
}

func Decode(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(dst); err != nil {
		return domain.Malformed()
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func Handle(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteError(w, r, err)
		}
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	var domainErr *domain.Error
	if errors.As(err, &domainErr) {
		write(w, apiError{
			Timestamp:  time.Now().UTC(),
			Status:     domainErr.Status(),
			Error:      domainErr.Title(),
			Message:    domainErr.Message,
			Path:       r.URL.Path,
			Violations: domainErr.Violations,
		})
		return
	}

	slog.Error("unhandled request error", "path", r.URL.Path, "error", err)
	write(w, apiError{
		Timestamp: time.Now().UTC(),
		Status:    http.StatusInternalServerError,
		Error:     "Internal Server Error",
		Message:   "An unexpected error occurred",
		Path:      r.URL.Path,
	})
}

func write(w http.ResponseWriter, payload apiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	_ = json.NewEncoder(w).Encode(payload)
}
