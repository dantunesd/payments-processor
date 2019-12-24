package payment

import (
	"context"
	"fmt"
)

// Payment succeeded
const (
	Authorized       = 1
	PaymentConfirmed = 2
)

// NewCieloStrategy strategy constructor
func NewCieloStrategy(r ICieloRepository) CieloStrategy {
	return CieloStrategy{
		r: r,
	}
}

// CieloStrategy .
type CieloStrategy struct {
	r ICieloRepository
}

// Process .
func (c CieloStrategy) Process(ctx context.Context, p Payment, s Source) error {
	crb := CieloRequestBody{
		MerchantOrderID: p.OrderID,
		Customer: Customer{
			Name: p.Customer.Name,
		},
		Payment: CieloRequestPayment{
			Type:         "CreditCard",
			Amount:       p.Details.Amount,
			Installments: p.Details.Installments,
			CreditCard: CieloCreditCard{
				CardNumber:     s.CardNumber,
				SecurityCode:   s.CVV,
				Holder:         p.Customer.Name,
				ExpirationDate: fmt.Sprintf("%d/%d", p.Details.Card.ExpirationMonth, p.Details.Card.ExpirationYear),
				Brand:          p.Details.Card.Brand,
			},
		},
	}
	res, err := c.r.Sale(ctx, crb)

	if err != nil {
		return NewInternalServerError("failed to comunicate with Cielo")
	}

	if !paymentSucceeded(res) {
		return NewPaymentError(res.Payment.ReturnMessage)
	}

	return nil
}

func paymentSucceeded(c *CieloResponseBody) bool {
	return c.Payment.Status == PaymentConfirmed || c.Payment.Status == Authorized
}
