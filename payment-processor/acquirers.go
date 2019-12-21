package payment

import "fmt"

// Consts for acquirer strategies
const (
	Cielo AcquirerStrategy = "Cielo"
	Rede  AcquirerStrategy = "Rede"
)

// AcquirerStrategy type for acquirer strategies
type AcquirerStrategy string

// IAcquirerStrategy interface for a AcquirerStrategy
type IAcquirerStrategy interface {
	Process(Payment, Source) error
}

// IAcquirerProvider interface for a AcquirerProvider
type IAcquirerProvider interface {
	GetAcquirer(AcquirerStrategy) IAcquirerStrategy
}

// Acquirers represents a list of acquirers
type Acquirers map[AcquirerStrategy]IAcquirerStrategy

// BuildAcquirers return active acquirers
func BuildAcquirers(cr ICieloRepository) Acquirers {
	return Acquirers{
		Cielo: NewCieloStrategy(cr),
		Rede:  NewRedeStrategy(),
	}
}

// AcquirerProvider is a acquirer provider
type AcquirerProvider struct {
	Acquirers Acquirers
}

// NewAcquirerProvider AcquirerProvider constructor
func NewAcquirerProvider(acquirers Acquirers) *AcquirerProvider {
	return &AcquirerProvider{acquirers}
}

// GetAcquirer returns a acquirer strategy
func (ap *AcquirerProvider) GetAcquirer(as AcquirerStrategy) IAcquirerStrategy {
	return ap.Acquirers[as]
}

// CieloStrategy .
type CieloStrategy struct {
	r ICieloRepository
}

// NewCieloStrategy .
func NewCieloStrategy(r ICieloRepository) CieloStrategy {
	return CieloStrategy{
		r: r,
	}
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

// RedeStrategy .
type RedeStrategy struct{}

// NewRedeStrategy .
func NewRedeStrategy() RedeStrategy {
	return RedeStrategy{}
}

// Process .
func (c RedeStrategy) Process(p Payment, s Source) error {
	fmt.Println("Processing Rede")
	return nil
}
