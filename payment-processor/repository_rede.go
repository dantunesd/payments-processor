package payment

import (
	"context"
)

const transactionPath = "/v1/transactions"

// IRedeRepository is a interface for Rede Repository.
type IRedeRepository interface {
	Transaction(context.Context, RedeRequestBody) (ITransaction, error)
}

// RedeRepository repository to comunicate with rede.
type RedeRepository struct {
	hr IHTTPRequester
}

// NewRedeRepository RedeRepository's constructor.
func NewRedeRepository(hr IHTTPRequester) IRedeRepository {
	return &RedeRepository{
		hr: hr,
	}
}

// Transaction makes a rede transaction.
func (c *RedeRepository) Transaction(ctx context.Context, rrb RedeRequestBody) (ITransaction, error) {

	r, err := c.hr.Post(ctx, transactionPath, rrb)
	if err != nil {
		return nil, err
	}

	return NewRedeTransaction(r), nil
}
