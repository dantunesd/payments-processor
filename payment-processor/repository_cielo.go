package payment

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

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

	r, err := c.r.Post(ctx, salePath, crb)

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
