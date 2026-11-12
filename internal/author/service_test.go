package author

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAutorName(t *testing.T) {
	repository := newFakeAuthorRepo()
	service := NewService(repository)

	response, err := service.Add(context.Background(), CreateAuthorRequest{
		Name:        "Mitch Albom",
		Biography:   "I am author of book called Tuesdays with morrie",
		Birth:       1958,
		Nationality: "American",
	})
	require.NoError(t, err)
	require.Equal(t, "Mitch Albom", response.Name)
	require.Equal(t, "I am author of book called Tuesdays with morrie", response.Biography)
	require.Equal(t, 1958, response.Birth)
	require.Equal(t, "American", response.Nationality)
}

type fakeAuthorRepo struct {
	byName map[string]Author
}

func newFakeAuthorRepo() *fakeAuthorRepo {
	return &fakeAuthorRepo{byName: map[string]Author{}}
}

func (f *fakeAuthorRepo) ExistByName(_ context.Context, name string) (bool, error) {
	_, ok := f.byName[name];
	return ok, nil
}

func (f *fakeAuthorRepo) Create(_ context.Context, author Author) (Author, error) {
	if _, ok := f.byName[author.Name]; ok {
		return Author{}, ErrDuplicate
	}
	author.Name = "Mitch Albom"
	f.byName[author.Name] = author
	return author, nil
}