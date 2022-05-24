package payment

import "net/http"

// ErrorType type for custom errors.
type ErrorType int

// ErrorTypes.
const (
	InvalidContentError ErrorType = iota
	InternalServerError
	TransactionError
)

// Error is a customized error
type Error struct {
	ErrorMessage string
	ErrorType    ErrorType
	StatusCode   int
}

// NewError Error's constructor.
func NewError(message string, errType ErrorType, statusCode int) *Error {
	return &Error{ErrorMessage: message, ErrorType: errType, StatusCode: statusCode}
}

// NewInvalidContentError InvalidContentError's constructor.
func NewInvalidContentError(message string) *Error {
	return NewError(message, InvalidContentError, http.StatusBadRequest)
}

// NewInternalServerError InternalServerError's constructor.
func NewInternalServerError(message string) *Error {
	return NewError(message, InternalServerError, http.StatusInternalServerError)
}

// NewTransactionError TransactionError's constructor.
func NewTransactionError(message string) *Error {
	return NewError(message, TransactionError, http.StatusBadRequest)
}

// Error return a customized error
func (e *Error) Error() string {
	return e.ErrorMessage
}
