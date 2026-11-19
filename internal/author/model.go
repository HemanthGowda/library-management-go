package author

import "time"

type Author struct {
	Id          string
	Name        string
	Biography   string
	Birth       int
	Nationality string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Biography   string `json:"biography"`
	Birth       int    `json:"birth"`
	Nationality string `json:"nationality"`
}

type AuthorResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Biography   string    `json:"biography"`
	Birth       int       `json:"birth"`
	Nationality string    `json:"nationality"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaginatedResponse struct {
	Data       []Author `json:"data"`
	Total      int      `json:"total"`
	Page       int      `json:"page"`
	TotalPages int      `json:"total pages"`
	Limit      int      `json:"limit"`
}
