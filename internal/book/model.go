package book

import "time"

type Book struct {
	ID              string
	Title           string
	AuthorName      string
	ISBN            string
	Description     string
	PublicationYear int
	TotalCopies     int
	AvailableCopies int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UpdateBookRequest struct {
	Title           *string `json:"title"`
	AuthorName      *string `json:"authorName"`
	ISBN            *string `json:"isbn"`
	Description     *string `json:"description"`
	PublicationYear *int    `json:"publicationYear"`
	TotalCopies     *int    `json:"totalCopies"`
	AvailableCopies *int    `json:"availableCopies"`
}

type BookResponse struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	AuthorName      string    `json:"authorName"`
	ISBN            string    `json:"isbn"`
	Description     string    `json:"description,omitempty"`
	PublicationYear int       `json:"publicationYear"`
	TotalCopies     int       `json:"totalCopies"`
	AvailableCopies int       `json:"availableCopies"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
