package payment

// IResponser interface for responser
type IResponser interface {
	GetStatusCode() int
	GetBody() []byte
}

// Response represents a response
type Response struct {
	StatusCode int
	Body       []byte
}

// NewResponse constructor
func NewResponse(statusCode int, body []byte) IResponser {
	return Response{
		StatusCode: statusCode,
		Body:       body,
	}
}

// GetStatusCode return status code
func (r Response) GetStatusCode() int {
	return r.StatusCode
}

// GetBody return body
func (r Response) GetBody() []byte {
	return r.Body
}
