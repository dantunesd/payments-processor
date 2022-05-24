package payment

import (
	"context"
	"fmt"
)

// CieloStrategy implementation of IAcquirerStrategy.
type CieloStrategy struct {
	cr ICieloRepository
}

// NewCieloStrategy CieloStrategy's constructor.
func NewCieloStrategy(cr ICieloRepository) *CieloStrategy {
	return &CieloStrategy{
		cr: cr,
	}
}

// Process processes the Cielo's transaction results.
func (c *CieloStrategy) Process(ctx context.Context, p Payment, s Source) error {

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

	t, err := c.cr.Transaction(ctx, crb)
	if err != nil {
		return err
	}

	return t.PaymentSucceeded()
}
