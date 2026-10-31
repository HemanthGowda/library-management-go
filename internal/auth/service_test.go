package auth

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/HemanthGowda/library-management-go/internal/domain"
	"github.com/HemanthGowda/library-management-go/internal/user"
)

func TestLoginReturnsBearerTokenForValidCredentials(t *testing.T) {
	repository := &memoryUsers{users: map[string]user.User{
		"admin@library.local": enabledAdmin(),
	}}
	service := newAuthService(t, repository)

	response, err := service.Login(context.Background(), LoginRequest{
		Email:    "admin@library.local",
		Password: "ChangeMe123!",
	})
	require.NoError(t, err)
	require.Equal(t, "Bearer", response.TokenType)
	require.NotEmpty(t, response.AccessToken)
	require.Equal(t, "admin@library.local", response.Email)
	require.Equal(t, []user.Role{user.RoleAdmin}, response.Roles)
	require.Equal(t, int64(time.Hour/time.Millisecond), response.ExpiresInMs)
}

func TestLoginRejectsUnknownUser(t *testing.T) {
	service := newAuthService(t, &memoryUsers{users: map[string]user.User{}})
	_, err := service.Login(context.Background(), LoginRequest{Email: "missing@library.local", Password: "Password1"})
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindUnauthorized, domainErr.Kind)
}

func TestLoginRejectsDisabledAccount(t *testing.T) {
	account := enabledAdmin()
	account.Enabled = false
	service := newAuthService(t, &memoryUsers{users: map[string]user.User{account.Email: account}})

	_, err := service.Login(context.Background(), LoginRequest{Email: account.Email, Password: "ChangeMe123!"})
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, "Account is disabled", domainErr.Message)
}

func TestLoginRejectsInvalidPassword(t *testing.T) {
	account := enabledAdmin()
	service := newAuthService(t, &memoryUsers{users: map[string]user.User{account.Email: account}})
	_, err := service.Login(context.Background(), LoginRequest{Email: account.Email, Password: "wrong"})
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindUnauthorized, domainErr.Kind)
}

func newAuthService(t *testing.T, repository user.Repository) *Service {
	t.Helper()
	tokens, err := NewJWTService("test-secret-key-that-is-long-enough-for-hs256-signing", "library-management-test", time.Hour)
	require.NoError(t, err)
	return NewService(repository, fakeEncoder{}, tokens)
}

func enabledAdmin() user.User {
	return user.User{
		ID:           "admin-1",
		Email:        "admin@library.local",
		PasswordHash: "hashed:ChangeMe123!",
		Roles:        []user.Role{user.RoleAdmin},
		Enabled:      true,
	}
}

type fakeEncoder struct{}

func (fakeEncoder) Encode(raw string) (string, error) { return "hashed:" + raw, nil }
func (fakeEncoder) Matches(raw, encoded string) bool  { return encoded == "hashed:"+raw }

type memoryUsers struct {
	users map[string]user.User
}

func (m *memoryUsers) ExistsByEmail(_ context.Context, email string) (bool, error) {
	_, ok := m.users[email]
	return ok, nil
}

func (m *memoryUsers) Create(_ context.Context, account user.User) (user.User, error) {
	m.users[account.Email] = account
	return account, nil
}

func (m *memoryUsers) FindByEmail(_ context.Context, email string) (user.User, error) {
	account, ok := m.users[email]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return account, nil
}
