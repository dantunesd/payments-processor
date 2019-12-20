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
	Process(Payment) error
}

// IAcquirerProvider interface for a AcquirerProvider
type IAcquirerProvider interface {
	GetAcquirer(AcquirerStrategy) IAcquirerStrategy
}

// Acquirers represents a list of acquirers
type Acquirers map[AcquirerStrategy]IAcquirerStrategy

// BuildAcquirers return active acquirers
func BuildAcquirers() Acquirers {
	return Acquirers{
		Cielo: NewCieloStrategy(),
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
type CieloStrategy struct{}

// NewCieloStrategy .
func NewCieloStrategy() CieloStrategy {
	return CieloStrategy{}
}

// Process .
func (c CieloStrategy) Process(p Payment) error {
	fmt.Println("Processing cielo")
	return nil
}

// RedeStrategy .
type RedeStrategy struct{}

// NewRedeStrategy .
func NewRedeStrategy() RedeStrategy {
	return RedeStrategy{}
}

// Process .
func (c RedeStrategy) Process(p Payment) error {
	fmt.Println("Processing Rede")
	return nil
}
