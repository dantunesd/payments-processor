package payment

import (
	"context"
	"fmt"
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
	transaction, err := c.r.Transaction(ctx, crb)
	if err != nil {
		return err
	}

	if !transaction.IsSucceeded() {
		return transaction.GetError()
	}

	return nil
}
