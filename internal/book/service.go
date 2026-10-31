package book

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/HemanthGowda/library-management-go/internal/domain"
)

var isbnPattern = regexp.MustCompile(`^(97[89])?\d{9}[\dXx]$`)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Update(ctx context.Context, bookID string, request UpdateBookRequest) (BookResponse, error) {
	if violations := validateUpdate(request); len(violations) > 0 {
		return BookResponse{}, domain.ValidationFailed(violations)
	}

	existing, err := s.repository.FindByID(ctx, bookID)
	if errors.Is(err, ErrNotFound) {
		return BookResponse{}, domain.NotFound("Book not found: " + bookID)
	}
	if err != nil {
		return BookResponse{}, err
	}

	isbn := normalizeISBN(*request.ISBN)
	taken, err := s.repository.ISBNUsedByOther(ctx, isbn, bookID)
	if err != nil {
		return BookResponse{}, err
	}
	if taken {
		return BookResponse{}, domain.Conflict("ISBN is already assigned to another book")
	}
	if *request.AvailableCopies > *request.TotalCopies {
		return BookResponse{}, domain.BadRequest("Available copies cannot exceed total copies")
	}

	existing.Title = strings.TrimSpace(*request.Title)
	existing.AuthorName = strings.TrimSpace(*request.AuthorName)
	existing.ISBN = isbn
	existing.Description = ""
	if request.Description != nil {
		existing.Description = strings.TrimSpace(*request.Description)
	}
	existing.PublicationYear = *request.PublicationYear
	existing.TotalCopies = *request.TotalCopies
	existing.AvailableCopies = *request.AvailableCopies

	updated, err := s.repository.Update(ctx, existing)
	if errors.Is(err, ErrNotFound) {
		return BookResponse{}, domain.NotFound("Book not found: " + bookID)
	}
	if errors.Is(err, ErrDuplicate) {
		return BookResponse{}, domain.Conflict("ISBN is already assigned to another book")
	}
	if err != nil {
		return BookResponse{}, err
	}
	return toResponse(updated), nil
}

func validateUpdate(request UpdateBookRequest) []domain.Violation {
	var violations []domain.Violation

	switch {
	case request.Title == nil || strings.TrimSpace(*request.Title) == "":
		violations = append(violations, domain.Violation{Field: "title", Message: "Title is required"})
	case utf8.RuneCountInString(strings.TrimSpace(*request.Title)) > 200:
		violations = append(violations, domain.Violation{Field: "title", Message: "Title must be between 1 and 200 characters"})
	}

	switch {
	case request.AuthorName == nil || strings.TrimSpace(*request.AuthorName) == "":
		violations = append(violations, domain.Violation{Field: "authorName", Message: "Author name is required"})
	case utf8.RuneCountInString(strings.TrimSpace(*request.AuthorName)) > 160:
		violations = append(violations, domain.Violation{Field: "authorName", Message: "Author name must be between 1 and 160 characters"})
	}

	switch {
	case request.ISBN == nil || strings.TrimSpace(*request.ISBN) == "":
		violations = append(violations, domain.Violation{Field: "isbn", Message: "ISBN is required"})
	case !isbnPattern.MatchString(strings.TrimSpace(*request.ISBN)):
		violations = append(violations, domain.Violation{Field: "isbn", Message: "ISBN must be a valid ISBN-10 or ISBN-13 without hyphens"})
	}

	if request.Description != nil && utf8.RuneCountInString(*request.Description) > 2000 {
		violations = append(violations, domain.Violation{Field: "description", Message: "Description must be at most 2000 characters"})
	}

	switch {
	case request.PublicationYear == nil:
		violations = append(violations, domain.Violation{Field: "publicationYear", Message: "Publication year is required"})
	case *request.PublicationYear < 1450:
		violations = append(violations, domain.Violation{Field: "publicationYear", Message: "Publication year must be 1450 or later"})
	case *request.PublicationYear > 2100:
		violations = append(violations, domain.Violation{Field: "publicationYear", Message: "Publication year must be 2100 or earlier"})
	}

	switch {
	case request.TotalCopies == nil:
		violations = append(violations, domain.Violation{Field: "totalCopies", Message: "Total copies is required"})
	case *request.TotalCopies < 1:
		violations = append(violations, domain.Violation{Field: "totalCopies", Message: "Total copies must be at least 1"})
	case *request.TotalCopies > 10000:
		violations = append(violations, domain.Violation{Field: "totalCopies", Message: "Total copies must be at most 10000"})
	}

	switch {
	case request.AvailableCopies == nil:
		violations = append(violations, domain.Violation{Field: "availableCopies", Message: "Available copies is required"})
	case *request.AvailableCopies < 0:
		violations = append(violations, domain.Violation{Field: "availableCopies", Message: "Available copies cannot be negative"})
	case *request.AvailableCopies > 10000:
		violations = append(violations, domain.Violation{Field: "availableCopies", Message: "Available copies must be at most 10000"})
	}

	return violations
}

func normalizeISBN(isbn string) string {
	return strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(isbn), "-", ""))
}

func toResponse(book Book) BookResponse {
	return BookResponse{
		ID:              book.ID,
		Title:           book.Title,
		AuthorName:      book.AuthorName,
		ISBN:            book.ISBN,
		Description:     book.Description,
		PublicationYear: book.PublicationYear,
		TotalCopies:     book.TotalCopies,
		AvailableCopies: book.AvailableCopies,
		CreatedAt:       book.CreatedAt.UTC(),
		UpdatedAt:       book.UpdatedAt.UTC(),
	}
}
