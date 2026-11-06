package user

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

func (h *Handler) Register(c fiber.Ctx) error {
	var request RegisterUserRequest
	if err := httpx.Decode(c, &request); err != nil {
		return err
	}

	response, err := h.service.Register(c.Context(), request)
	if err != nil {
		return err
	}
	return httpx.WriteJSON(c, fiber.StatusCreated, response)
}
