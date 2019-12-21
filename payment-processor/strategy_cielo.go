package payment

import (
	"fmt"
	"time"
)

// NewCieloStrategy strategy constructor
func NewCieloStrategy(baseURI, merchantID, merchantKey string, timeout time.Duration) CieloStrategy {
	return CieloStrategy{
		r: NewCieloRepository(
			NewHTTPRequester(
				baseURI,
				headers{
					"merchantid":  merchantID,
					"merchantkey": merchantKey,
				},
				timeout,
			),
		),
	}
}

// CieloStrategy .
type CieloStrategy struct {
	r ICieloRepository
}

// Process .
func (c CieloStrategy) Process(p Payment, s Source) error {
	crb := CieloRequestBody{
		MerchantOrderID: "newOrder",
		Customer: Customer{
			Name: p.Customer.Name,
		},
		Payment: CieloPayment{
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
	res, err := c.r.Sale(crb)

	fmt.Println("Processing Cielo", res)

	return err
}
