package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/HemanthGowda/library-management-go/internal/user"
)

type JWTService struct {
	secret     []byte
	issuer     string
	expiration time.Duration
}

type claims struct {
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

func NewJWTService(secret, issuer string, expiration time.Duration) (*JWTService, error) {
	if len(secret) < 32 {
		return nil, fmt.Errorf("JWT secret must be at least 256 bits (32 bytes)")
	}
	return &JWTService{secret: []byte(secret), issuer: issuer, expiration: expiration}, nil
}

func (s *JWTService) Expiration() time.Duration {
	return s.expiration
}

func (s *JWTService) Generate(account user.User) (string, error) {
	now := time.Now().UTC()
	roles := make([]string, 0, len(account.Roles))
	for _, role := range account.Roles {
		roles = append(roles, string(role))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Email: account.Email,
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   account.ID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expiration)),
		},
	})
	return token.SignedString(s.secret)
}

func (s *JWTService) Parse(token string) (Principal, error) {
	parsed, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.secret, nil
	}, jwt.WithIssuer(s.issuer), jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return Principal{}, err
	}

	payload, ok := parsed.Claims.(*claims)
	if !ok || !parsed.Valid {
		return Principal{}, fmt.Errorf("invalid token")
	}

	roles := make([]user.Role, 0, len(payload.Roles))
	for _, role := range payload.Roles {
		roles = append(roles, user.Role(role))
	}
	return Principal{UserID: payload.Subject, Email: payload.Email, Roles: roles}, nil
}
