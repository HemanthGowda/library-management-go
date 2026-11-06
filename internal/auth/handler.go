package auth

import (
	"github.com/gofiber/fiber/v3"

	"github.com/HemanthGowda/library-management-go/internal/httpx"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Login(c fiber.Ctx) error {
	var request LoginRequest
	if err := httpx.Decode(c, &request); err != nil {
		return err
	}

	response, err := h.service.Login(c.Context(), request)
	if err != nil {
		return err
	}
	return httpx.WriteJSON(c, fiber.StatusOK, response)
}
