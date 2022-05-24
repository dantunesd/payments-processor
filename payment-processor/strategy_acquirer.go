package payment

import "context"

// Available acquirers.
const (
	Cielo Acquirer = "Cielo"
	Rede  Acquirer = "Rede"
)

// Acquirer type
type Acquirer string

// AcquirerStrategies is map for acquirers and strategies.
type AcquirerStrategies map[Acquirer]IAcquirerStrategy

// IAcquirerStrategy interface for acquirer strategy.
type IAcquirerStrategy interface {
	Process(context.Context, Payment, Source) error
}

// IAcquirerProvider interface for acquirer provider.
type IAcquirerProvider interface {
	GetAcquirerStrategy(Acquirer) IAcquirerStrategy
}

// AcquirerProvider implementation of IAcquirerProvider.
type AcquirerProvider struct {
	As AcquirerStrategies
}

// NewAcquirerProvider AcquirerProvider's constructor.
func NewAcquirerProvider(as AcquirerStrategies) *AcquirerProvider {
	return &AcquirerProvider{as}
}

// GetAcquirerStrategy returns a acquirer strategy.
func (a *AcquirerProvider) GetAcquirerStrategy(acquirer Acquirer) IAcquirerStrategy {
	return a.As[acquirer]
}
