package payment

import (
	"context"
)

const transactionPath = "/v1/transactions"

// IRedeRepository is a interface for Rede Repository
type IRedeRepository interface {
	Transaction(context.Context, RedeRequestBody) (ITransaction, error)
}

// RedeRepository repository to comunicate with cielo
type RedeRepository struct {
	r IHTTPRequester
}

// NewRedeRepository constructor
func NewRedeRepository(r IHTTPRequester) IRedeRepository {
	return &RedeRepository{
		r: r,
	}
}

// Transaction makes a rede transaction
func (c *RedeRepository) Transaction(ctx context.Context, rrb RedeRequestBody) (ITransaction, error) {

	r, err := c.r.Post(ctx, transactionPath, rrb)
	if err != nil {
		return nil, err
	}

	return NewRedeTransaction(r), nil
}
