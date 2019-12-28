package payment

import (
	"context"
	"fmt"
)

// RedeStrategy implementation of IAcquirerStrategy.
type RedeStrategy struct {
	rr IRedeRepository
}

// NewRedeStrategy RedeStrategy's constructor.
func NewRedeStrategy(rr IRedeRepository) *RedeStrategy {
	return &RedeStrategy{
		rr: rr,
	}
}

// Process processes the Rede's transaction results.
func (r *RedeStrategy) Process(ctx context.Context, p Payment, s Source) error {

	rrb := RedeRequestBody{
		Capture:         true,
		Reference:       p.OrderID,
		Amount:          p.Details.Amount,
		HolderName:      p.Customer.Name,
		CardNumber:      s.CardNumber,
		Installments:    p.Details.Installments,
		ExpirationMonth: p.Details.Card.ExpirationMonth,
		ExpirationYear:  p.Details.Card.ExpirationYear,
		SecurityCode:    fmt.Sprintf("%d", s.CVV),
	}

	t, err := r.rr.Transaction(ctx, rrb)
	if err != nil {
		return err
	}

	return t.PaymentSucceeded()
}
