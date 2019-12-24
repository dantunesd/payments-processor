package payment

import "context"

const salePath = "/1/sales"

// ICieloRepository is a interface for Cielo Repository
type ICieloRepository interface {
	Sale(context.Context, CieloRequestBody) (*CieloResponseBody, error)
}

// CieloRepository repository to comunicate with cielo
type CieloRepository struct {
	r IHTTPRequester
}

// NewCieloRepository constructor
func NewCieloRepository(r IHTTPRequester) ICieloRepository {
	return &CieloRepository{
		r: r,
	}
}

// Sale .
func (c *CieloRepository) Sale(ctx context.Context, crb CieloRequestBody) (*CieloResponseBody, error) {
	out := &CieloResponseBody{}
	if err := c.r.Post(ctx, salePath, crb, out); err != nil {
		return out, err
	}
	return out, nil
}
