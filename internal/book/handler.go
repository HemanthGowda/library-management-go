package book

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/HemanthGowda/library-management-go/internal/auth"
	"github.com/HemanthGowda/library-management-go/internal/domain"
	"github.com/HemanthGowda/library-management-go/internal/httpx"
	"github.com/HemanthGowda/library-management-go/internal/user"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) error {
	principal, ok := auth.PrincipalFrom(r.Context())
	if !ok {
		return domain.Unauthorized("Authentication is required")
	}
	if !principal.HasAny(user.RoleAdmin, user.RoleLibrarian) {
		return domain.Forbidden()
	}

	var request UpdateBookRequest
	if err := httpx.Decode(r, &request); err != nil {
		return err
	}

	response, err := h.service.Update(r.Context(), chi.URLParam(r, "id"), request)
	if err != nil {
		return err
	}
	return httpx.WriteJSON(w, http.StatusOK, response)
}
