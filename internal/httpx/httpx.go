package httpx

import (
	"errors"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"

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

func Decode(c fiber.Ctx, dst any) error {
	if err := c.Bind().JSON(dst); err != nil {
		return domain.Malformed()
	}
	return nil
}

func WriteJSON(c fiber.Ctx, status int, payload any) error {
	return c.Status(status).JSON(payload)
}

func WriteError(c fiber.Ctx, err error) error {
	var domainErr *domain.Error
	if errors.As(err, &domainErr) {
		return write(c, apiError{
			Timestamp:  time.Now().UTC(),
			Status:     domainErr.Status(),
			Error:      domainErr.Title(),
			Message:    domainErr.Message,
			Path:       c.Path(),
			Violations: domainErr.Violations,
		})
	}

	slog.Error("unhandled request error", "path", c.Path(), "error", err)
	return write(c, apiError{
		Timestamp: time.Now().UTC(),
		Status:    fiber.StatusInternalServerError,
		Error:     "Internal Server Error",
		Message:   "An unexpected error occurred",
		Path:      c.Path(),
	})
}

func write(c fiber.Ctx, payload apiError) error {
	return c.Status(payload.Status).JSON(payload)
}
