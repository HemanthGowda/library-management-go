package user

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/HemanthGowda/library-management-go/internal/domain"
)

func TestRegisterCreatesMemberWithNormalizedEmail(t *testing.T) {
	repository := newFakeUserRepo()
	service := NewService(repository, fakeEncoder{})

	response, err := service.Register(context.Background(), RegisterUserRequest{
		Email:    "  Alex@Example.com ",
		Password: "Password1",
		FullName: " Alex Reader ",
	})
	require.NoError(t, err)
	require.Equal(t, "alex@example.com", response.Email)
	require.Equal(t, "Alex Reader", response.FullName)
	require.Equal(t, []Role{RoleMember}, response.Roles)
	require.True(t, response.Enabled)
	require.Equal(t, "hashed:Password1", repository.byEmail["alex@example.com"].PasswordHash)
}

func TestRegisterRejectsDuplicateEmail(t *testing.T) {
	repository := newFakeUserRepo()
	repository.byEmail["alex@example.com"] = User{Email: "alex@example.com"}
	service := NewService(repository, fakeEncoder{})

	_, err := service.Register(context.Background(), RegisterUserRequest{
		Email:    "alex@example.com",
		Password: "Password1",
		FullName: "Alex Reader",
	})
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindConflict, domainErr.Kind)
	require.Contains(t, domainErr.Message, "already registered")
}

type fakeEncoder struct{}

func (fakeEncoder) Encode(raw string) (string, error) { return "hashed:" + raw, nil }
func (fakeEncoder) Matches(raw, encoded string) bool  { return encoded == "hashed:"+raw }

type fakeUserRepo struct {
	byEmail map[string]User
}

func newFakeUserRepo() *fakeUserRepo {
	return &fakeUserRepo{byEmail: map[string]User{}}
}

func (f *fakeUserRepo) ExistsByEmail(_ context.Context, email string) (bool, error) {
	_, ok := f.byEmail[email]
	return ok, nil
}

func (f *fakeUserRepo) Create(_ context.Context, user User) (User, error) {
	if _, ok := f.byEmail[user.Email]; ok {
		return User{}, ErrDuplicate
	}
	user.ID = "user-1"
	user.CreatedAt = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	f.byEmail[user.Email] = user
	return user, nil
}

func (f *fakeUserRepo) FindByEmail(_ context.Context, email string) (User, error) {
	user, ok := f.byEmail[email]
	if !ok {
		return User{}, ErrNotFound
	}
	return user, nil
}
