package payment

// ErrorType error types.
type ErrorType int

// Error types values.
const (
	InvalidContent ErrorType = iota
	InvalidRequest
	PaymentFailed
)

// Error is a representation for error
type Error struct {
	ErrorMessage string
	ErrorType    ErrorType
}

// NewError constructor of Error.
func NewError(message string, errType ErrorType) *Error {
	return &Error{ErrorMessage: message, ErrorType: errType}
}

// NewInvalidContentError constructor of invalid content error.
func NewInvalidContentError(message string) *Error {
	return NewError(message, InvalidContent)
}

// NewInvalidRequestError constructor of invalid request to acquirer.
func NewInvalidRequestError(message string) *Error {
	return NewError(message, InvalidRequest)
}

// NewPaymentFailedError constructor of fail to pay.
func NewPaymentFailedError(message string) *Error {
	return NewError(message, PaymentFailed)
}

// Error returns error message
func (e *Error) Error() string {
	return e.ErrorMessage
}
