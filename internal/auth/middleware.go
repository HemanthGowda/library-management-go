package auth

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Middleware(tokens *JWTService) fiber.Handler {
	return func(c fiber.Ctx) error {
		header := c.Get(fiber.HeaderAuthorization)
		if !strings.HasPrefix(header, "Bearer ") {
			return c.Next()
		}

		principal, err := tokens.Parse(strings.TrimSpace(strings.TrimPrefix(header, "Bearer ")))
		if err != nil {
			return c.Next()
		}

		c.SetContext(WithPrincipal(c.Context(), principal))
		return c.Next()
	}
}
