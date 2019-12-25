package payment

import (
	"context"
	"fmt"
)

// RedeStrategy .
type RedeStrategy struct {
	r IRedeRepository
}

// NewRedeStrategy .
func NewRedeStrategy(r IRedeRepository) RedeStrategy {
	return RedeStrategy{
		r: r,
	}
}

// Process .
func (r RedeStrategy) Process(ctx context.Context, p Payment, s Source) error {

	rrb := RedeRequestBody{
		Capture:         true,
		Reference:       p.OrderID,
		Amount:          p.Details.Amount,
		HolderName:      p.Customer.Name,
		CardNumber:      s.CardNumber,
		ExpirationMonth: p.Details.Card.ExpirationMonth,
		ExpirationYear:  p.Details.Card.ExpirationYear,
		SecurityCode:    fmt.Sprintf("%d", s.CVV),
	}

	transaction, err := r.r.Transaction(ctx, rrb)
	if err != nil {
		return err
	}

	return transaction.PaymentSucceeded()
}
