package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/HemanthGowda/library-management-go/internal/user"
)

func TestGenerateAndParseTokenRoundTrip(t *testing.T) {
	service, err := NewJWTService("test-secret-key-that-is-long-enough-for-hs256-signing", "library-management-test", time.Hour)
	require.NoError(t, err)

	token, err := service.Generate(user.User{
		ID:    "user-42",
		Email: "admin@library.local",
		Roles: []user.Role{user.RoleAdmin, user.RoleLibrarian},
	})
	require.NoError(t, err)

	principal, err := service.Parse(token)
	require.NoError(t, err)
	require.Equal(t, "user-42", principal.UserID)
	require.Equal(t, "admin@library.local", principal.Email)
	require.ElementsMatch(t, []user.Role{user.RoleAdmin, user.RoleLibrarian}, principal.Roles)
	require.Equal(t, time.Hour, service.Expiration())
}

func TestRejectsShortSecret(t *testing.T) {
	_, err := NewJWTService("too-short", "issuer", time.Second)
	require.Error(t, err)
	require.Contains(t, err.Error(), "256 bits")
}
