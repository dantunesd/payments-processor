package payment

import "net/http"

// ErrorType error types.
type ErrorType int

// Error types values.
const (
	InvalidContentError ErrorType = iota
	InternalServerError
	IntegrationError
	EmissorError
)

// Error is a representation for error
type Error struct {
	ErrorMessage string
	ErrorType    ErrorType
	StatusCode   int
}

// NewError constructor of Error.
func NewError(message string, errType ErrorType, statusCode int) *Error {
	return &Error{ErrorMessage: message, ErrorType: errType, StatusCode: statusCode}
}

// NewInvalidContentError constructor of InvalidContentError.
func NewInvalidContentError(message string) *Error {
	return NewError(message, InvalidContentError, http.StatusBadRequest)
}

// NewInternalServerError constructor of InternalServerError.
func NewInternalServerError(message string) *Error {
	return NewError(message, InternalServerError, http.StatusInternalServerError)
}

// NewIntegrationError constructor of IntegrationError.
func NewIntegrationError(message string) *Error {
	return NewError(message, IntegrationError, http.StatusBadRequest)
}

// NewEmissorError constructor of EmissorError.
func NewEmissorError(message string) *Error {
	return NewError(message, EmissorError, http.StatusBadRequest)
}

// Error returns error message
func (e *Error) Error() string {
	return e.ErrorMessage
}
