package book

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/HemanthGowda/library-management-go/internal/domain"
)

func TestUpdateReplacesMutableFields(t *testing.T) {
	repository := newFakeBookRepo()
	repository.books["book-1"] = Book{ID: "book-1", Title: "Old Title", ISBN: "9780134685991", TotalCopies: 2, AvailableCopies: 1}
	service := NewService(repository)

	response, err := service.Update(context.Background(), "book-1", UpdateBookRequest{
		Title:           str("New Title"),
		AuthorName:      str("New Author"),
		ISBN:            str("9780596009205"),
		Description:     str("Updated description"),
		PublicationYear: integer(2020),
		TotalCopies:     integer(5),
		AvailableCopies: integer(4),
	})
	require.NoError(t, err)
	require.Equal(t, "New Title", response.Title)
	require.Equal(t, "New Author", response.AuthorName)
	require.Equal(t, "9780596009205", response.ISBN)
	require.Equal(t, 5, response.TotalCopies)
	require.Equal(t, 4, response.AvailableCopies)
}

func TestUpdateFailsWhenBookMissing(t *testing.T) {
	service := NewService(newFakeBookRepo())
	_, err := service.Update(context.Background(), "missing", sampleRequest())
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindNotFound, domainErr.Kind)
}

func TestUpdateFailsWhenISBNConflicts(t *testing.T) {
	repository := newFakeBookRepo()
	repository.books["book-1"] = Book{ID: "book-1", ISBN: "9780134685991"}
	repository.books["book-2"] = Book{ID: "book-2", ISBN: "9780596009205"}
	service := NewService(repository)

	_, err := service.Update(context.Background(), "book-1", sampleRequest())
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindConflict, domainErr.Kind)
	require.Contains(t, domainErr.Message, "ISBN")
}

func TestUpdateFailsWhenAvailableExceedsTotal(t *testing.T) {
	repository := newFakeBookRepo()
	repository.books["book-1"] = Book{ID: "book-1", ISBN: "9780134685991"}
	service := NewService(repository)

	_, err := service.Update(context.Background(), "book-1", UpdateBookRequest{
		Title:           str("Title"),
		AuthorName:      str("Author"),
		ISBN:            str("9780596009205"),
		Description:     str("Desc"),
		PublicationYear: integer(2020),
		TotalCopies:     integer(2),
		AvailableCopies: integer(3),
	})
	var domainErr *domain.Error
	require.ErrorAs(t, err, &domainErr)
	require.Equal(t, domain.KindBadRequest, domainErr.Kind)
	require.Contains(t, domainErr.Message, "Available copies")
}

func sampleRequest() UpdateBookRequest {
	return UpdateBookRequest{
		Title:           str("Title"),
		AuthorName:      str("Author"),
		ISBN:            str("9780596009205"),
		Description:     str("Desc"),
		PublicationYear: integer(2020),
		TotalCopies:     integer(5),
		AvailableCopies: integer(4),
	}
}

func str(value string) *string { return &value }
func integer(value int) *int   { return &value }

type fakeBookRepo struct {
	books map[string]Book
}

func newFakeBookRepo() *fakeBookRepo {
	return &fakeBookRepo{books: map[string]Book{}}
}

func (f *fakeBookRepo) FindByID(_ context.Context, id string) (Book, error) {
	book, ok := f.books[id]
	if !ok {
		return Book{}, ErrNotFound
	}
	return book, nil
}

func (f *fakeBookRepo) ISBNUsedByOther(_ context.Context, isbn, id string) (bool, error) {
	for _, book := range f.books {
		if book.ISBN == isbn && book.ID != id {
			return true, nil
		}
	}
	return false, nil
}

func (f *fakeBookRepo) Update(_ context.Context, book Book) (Book, error) {
	if _, ok := f.books[book.ID]; !ok {
		return Book{}, ErrNotFound
	}
	f.books[book.ID] = book
	return book, nil
}

func (f *fakeBookRepo) ExistsByISBN(_ context.Context, isbn string) (bool, error) {
	for _, book := range f.books {
		if book.ISBN == isbn {
			return true, nil
		}
	}
	return false, nil
}

func (f *fakeBookRepo) Insert(_ context.Context, book Book) (Book, error) {
	f.books[book.ID] = book
	return book, nil
}
