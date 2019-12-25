package payment

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

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

	r, err := c.r.Post(ctx, transactionPath, rrb)
	if err != nil {
		return out, err
	}

	if r.GetStatusCode() < http.StatusOK || r.GetStatusCode() >= http.StatusMultipleChoices {
		return out, errors.New(string(r.GetBody()))
	}

	if uErr := json.Unmarshal(r.GetBody(), out); uErr != nil {
		return out, uErr
	}

	return out, nil
}
