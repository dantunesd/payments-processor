package payment

import "context"

const transactionPath = "/v1/transactions"

// IRedeRepository is a interface for Rede Repository
type IRedeRepository interface {
	Transaction(context.Context, RedeRequestBody) (*RedeResponseBody, error)
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

// Transaction .
func (c *RedeRepository) Transaction(ctx context.Context, rrb RedeRequestBody) (*RedeResponseBody, error) {
	out := &RedeResponseBody{}

	if err := c.r.Post(ctx, transactionPath, rrb, out); err != nil {
		return out, err
	}
	return out, nil
}
