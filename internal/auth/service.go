package auth

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/HemanthGowda/library-management-go/internal/domain"
	"github.com/HemanthGowda/library-management-go/internal/password"
	"github.com/HemanthGowda/library-management-go/internal/user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string      `json:"accessToken"`
	TokenType   string      `json:"tokenType"`
	ExpiresInMs int64       `json:"expiresInMs"`
	UserID      string      `json:"userId"`
	Email       string      `json:"email"`
	Roles       []user.Role `json:"roles"`
}

type Service struct {
	users     user.Repository
	passwords password.Encoder
	tokens    *JWTService
}

func NewService(users user.Repository, passwords password.Encoder, tokens *JWTService) *Service {
	return &Service{users: users, passwords: passwords, tokens: tokens}
}

func (s *Service) Login(ctx context.Context, request LoginRequest) (LoginResponse, error) {
	email := strings.ToLower(strings.TrimSpace(request.Email))
	if violations := validateLogin(email, request.Password); len(violations) > 0 {
		return LoginResponse{}, domain.ValidationFailed(violations)
	}

	account, err := s.users.FindByEmail(ctx, email)
	if errors.Is(err, user.ErrNotFound) {
		return LoginResponse{}, domain.Unauthorized("Invalid email or password")
	}
	if err != nil {
		return LoginResponse{}, err
	}
	if !account.Enabled {
		return LoginResponse{}, domain.Unauthorized("Account is disabled")
	}
	if !s.passwords.Matches(request.Password, account.PasswordHash) {
		return LoginResponse{}, domain.Unauthorized("Invalid email or password")
	}

	token, err := s.tokens.Generate(account)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresInMs: s.tokens.Expiration().Milliseconds(),
		UserID:      account.ID,
		Email:       account.Email,
		Roles:       account.Roles,
	}, nil
}

func validateLogin(email, rawPassword string) []domain.Violation {
	var violations []domain.Violation
	if email == "" {
		violations = append(violations, domain.Violation{Field: "email", Message: "Email is required"})
	} else if !validEmail(email) {
		violations = append(violations, domain.Violation{Field: "email", Message: "Email must be valid"})
	}
	if rawPassword == "" {
		violations = append(violations, domain.Violation{Field: "password", Message: "Password is required"})
	}
	return violations
}

func validEmail(email string) bool {
	address, err := mail.ParseAddress(email)
	if err != nil || address.Address != email {
		return false
	}
	parts := strings.Split(email, "@")
	return len(parts) == 2 && strings.Contains(parts[1], ".")
}
