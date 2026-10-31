package user

import (
	"net/http"

	"github.com/HemanthGowda/library-management-go/internal/httpx"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) error {
	var request RegisterUserRequest
	if err := httpx.Decode(r, &request); err != nil {
		return err
	}

	response, err := h.service.Register(r.Context(), request)
	if err != nil {
		return err
	}
	return httpx.WriteJSON(w, http.StatusCreated, response)
}
