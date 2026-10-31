package user

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"

	"github.com/HemanthGowda/library-management-go/internal/domain"
	"github.com/HemanthGowda/library-management-go/internal/password"
)

type Service struct {
	repository Repository
	passwords  password.Encoder
}

func NewService(repository Repository, passwords password.Encoder) *Service {
	return &Service{repository: repository, passwords: passwords}
}

func (s *Service) Register(ctx context.Context, request RegisterUserRequest) (UserResponse, error) {
	email := strings.ToLower(strings.TrimSpace(request.Email))
	fullName := strings.TrimSpace(request.FullName)
	if violations := validateRegister(email, request.Password, fullName); len(violations) > 0 {
		return UserResponse{}, domain.ValidationFailed(violations)
	}

	exists, err := s.repository.ExistsByEmail(ctx, email)
	if err != nil {
		return UserResponse{}, err
	}
	if exists {
		return UserResponse{}, domain.Conflict("Email is already registered")
	}

	hash, err := s.passwords.Encode(request.Password)
	if err != nil {
		return UserResponse{}, err
	}

	created, err := s.repository.Create(ctx, User{
		Email:        email,
		PasswordHash: hash,
		FullName:     fullName,
		Roles:        []Role{RoleMember},
		Enabled:      true,
	})
	if errors.Is(err, ErrDuplicate) {
		return UserResponse{}, domain.Conflict("Email is already registered")
	}
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(created), nil
}

func validateRegister(email, rawPassword, fullName string) []domain.Violation {
	var violations []domain.Violation
	switch {
	case email == "":
		violations = append(violations, domain.Violation{Field: "email", Message: "Email is required"})
	case utf8.RuneCountInString(email) > 255:
		violations = append(violations, domain.Violation{Field: "email", Message: "Email must be at most 255 characters"})
	case !validEmail(email):
		violations = append(violations, domain.Violation{Field: "email", Message: "Email must be valid"})
	}

	switch {
	case rawPassword == "":
		violations = append(violations, domain.Violation{Field: "password", Message: "Password is required"})
	case utf8.RuneCountInString(rawPassword) < 8 || utf8.RuneCountInString(rawPassword) > 72:
		violations = append(violations, domain.Violation{Field: "password", Message: "Password must be between 8 and 72 characters"})
	case !containsLetterAndNumber(rawPassword):
		violations = append(violations, domain.Violation{Field: "password", Message: "Password must contain at least one letter and one number"})
	}

	switch {
	case fullName == "":
		violations = append(violations, domain.Violation{Field: "fullName", Message: "Full name is required"})
	case utf8.RuneCountInString(fullName) < 2 || utf8.RuneCountInString(fullName) > 120:
		violations = append(violations, domain.Violation{Field: "fullName", Message: "Full name must be between 2 and 120 characters"})
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

func containsLetterAndNumber(value string) bool {
	var letter, digit bool
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
			letter = true
		case r >= '0' && r <= '9':
			digit = true
		}
	}
	return letter && digit
}

func toResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Roles:     user.Roles,
		Enabled:   user.Enabled,
		CreatedAt: user.CreatedAt.UTC(),
	}
}
