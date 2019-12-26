package payment

import "context"

// Acquirer type for acquirers
type Acquirer string

// Consts for acquirers
const (
	Cielo Acquirer = "Cielo"
	Rede  Acquirer = "Rede"
)

// AcquirersStrategy map for acquirer and strategy
type AcquirersStrategy map[Acquirer]IAcquirerStrategy

// IAcquirerStrategy interface for a AcquirerStrategy
type IAcquirerStrategy interface {
	Process(context.Context, Payment, Source) error
}

// IAcquirerProvider interface for a AcquirerProvider
type IAcquirerProvider interface {
	GetAcquirerStrategy(Acquirer) IAcquirerStrategy
}

// AcquirerProvider is a acquirer provider
type AcquirerProvider struct {
	As AcquirersStrategy
}

// NewAcquirerProvider AcquirerProvider constructor
func NewAcquirerProvider(as AcquirersStrategy) *AcquirerProvider {
	return &AcquirerProvider{as}
}

// GetAcquirerStrategy returns a acquirer strategy
func (a *AcquirerProvider) GetAcquirerStrategy(acquirer Acquirer) IAcquirerStrategy {
	return a.As[acquirer]
}
