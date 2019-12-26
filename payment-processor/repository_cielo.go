package payment

import (
	"context"
)

const salePath = "/1/sales"

// ICieloRepository is a interface for Cielo Repository.
type ICieloRepository interface {
	Transaction(context.Context, CieloRequestBody) (ITransaction, error)
}

// CieloRepository repository to comunicate with cielo.
type CieloRepository struct {
	hr IHTTPRequester
}

// NewCieloRepository CieloRepository's constructor
func NewCieloRepository(hr IHTTPRequester) *CieloRepository {
	return &CieloRepository{
		hr: hr,
	}
}

// Transaction makes a cielo transaction.
func (c *CieloRepository) Transaction(ctx context.Context, crb CieloRequestBody) (ITransaction, error) {

	r, err := c.hr.Post(ctx, salePath, crb)

	if err != nil {
		return nil, err
	}

	return NewCieloTransaction(r), nil
}
