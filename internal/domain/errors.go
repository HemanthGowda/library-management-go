package domain

import "net/http"

type Kind int

const (
	KindBadRequest Kind = iota
	KindValidation
	KindMalformed
	KindUnauthorized
	KindForbidden
	KindNotFound
	KindConflict
)

type Violation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Kind       Kind
	Message    string
	Violations []Violation
}

func (e *Error) Error() string {
	if e.Message == "" {
		return http.StatusText(e.Status())
	}
	return e.Message
}

func (e *Error) Status() int {
	switch e.Kind {
	case KindValidation, KindBadRequest, KindMalformed:
		return http.StatusBadRequest
	case KindUnauthorized:
		return http.StatusUnauthorized
	case KindForbidden:
		return http.StatusForbidden
	case KindNotFound:
		return http.StatusNotFound
	case KindConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (e *Error) Title() string {
	switch e.Kind {
	case KindValidation:
		return "Validation Failed"
	case KindMalformed:
		return "Malformed JSON"
	case KindBadRequest:
		return "Bad Request"
	case KindUnauthorized:
		return "Unauthorized"
	case KindForbidden:
		return "Forbidden"
	case KindNotFound:
		return "Not Found"
	case KindConflict:
		return "Conflict"
	default:
		return "Internal Server Error"
	}
}

func BadRequest(message string) *Error {
	return &Error{Kind: KindBadRequest, Message: message}
}

func Malformed() *Error {
	return &Error{Kind: KindMalformed, Message: "Request body could not be parsed"}
}

func ValidationFailed(violations []Violation) *Error {
	return &Error{
		Kind:       KindValidation,
		Message:    "Request validation failed",
		Violations: violations,
	}
}

func Unauthorized(message string) *Error {
	return &Error{Kind: KindUnauthorized, Message: message}
}

func Forbidden() *Error {
	return &Error{Kind: KindForbidden, Message: "You do not have permission to perform this action"}
}

func NotFound(message string) *Error {
	return &Error{Kind: KindNotFound, Message: message}
}

func Conflict(message string) *Error {
	return &Error{Kind: KindConflict, Message: message}
}
