package auth

import (
	"context"

	"github.com/HemanthGowda/library-management-go/internal/user"
)

type principalKey struct{}

type Principal struct {
	UserID string
	Email  string
	Roles  []user.Role
}

func (p Principal) HasAny(roles ...user.Role) bool {
	for _, candidate := range roles {
		for _, owned := range p.Roles {
			if owned == candidate {
				return true
			}
		}
	}
	return false
}

func WithPrincipal(ctx context.Context, principal Principal) context.Context {
	return context.WithValue(ctx, principalKey{}, principal)
}

func PrincipalFrom(ctx context.Context) (Principal, bool) {
	principal, ok := ctx.Value(principalKey{}).(Principal)
	return principal, ok
}
