package payment

import (
	"context"
	"fmt"
)

// RedePaymentConfirmed succeeded
const RedePaymentConfirmed = "00"

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

	res, err := r.r.Transaction(ctx, rrb)
	if err != nil {
		return NewIntegrationError(("failed to comunicate with Rede"))
	}

	if !r.paymentSucceeded(res) {
		return NewTransactionError(res.ReturnMessage)
	}

	return nil
}

func (r RedeStrategy) paymentSucceeded(rrb *RedeResponseBody) bool {
	return rrb.ReturnCode == RedePaymentConfirmed
}
