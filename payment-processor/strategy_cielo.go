package payment

import (
	"context"
	"fmt"
)

// NewCieloStrategy strategy constructor
func NewCieloStrategy(cr ICieloRepository) CieloStrategy {
	return CieloStrategy{
		cr: cr,
	}
}

// CieloStrategy .
type CieloStrategy struct {
	cr ICieloRepository
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

	t, err := c.cr.Transaction(ctx, crb)
	if err != nil {
		return err
	}

	return t.PaymentSucceeded()
}
