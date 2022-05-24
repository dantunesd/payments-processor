package payment

// IResponser interface for http response.
type IResponser interface {
	GetStatusCode() int
	GetBody() []byte
}

// Response represents a http response.
type Response struct {
	StatusCode int
	Body       []byte
}

// NewResponse Response's constructor.
func NewResponse(statusCode int, body []byte) *Response {
	return &Response{
		StatusCode: statusCode,
		Body:       body,
	}
}

// GetStatusCode returns the status code.
func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// GetBody returns the body content.
func (r *Response) GetBody() []byte {
	return r.Body
}
