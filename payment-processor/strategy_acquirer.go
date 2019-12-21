package payment

// Consts for acquirers
const (
	Cielo Acquirer = "Cielo"
	Rede  Acquirer = "Rede"
)

// Acquirer type for acquirers
type Acquirer string

// AcquirersStrategy map for acquirer and strategy
type AcquirersStrategy map[Acquirer]IAcquirerStrategy

// IAcquirerStrategy interface for a AcquirerStrategy
type IAcquirerStrategy interface {
	Process(Payment, Source) error
}

// IAcquirerProvider interface for a AcquirerProvider
type IAcquirerProvider interface {
	GetAcquirerStrategy(Acquirer) IAcquirerStrategy
}

// AcquirerProvider is a acquirer provider
type AcquirerProvider struct {
	Acquirers AcquirersStrategy
}

// NewAcquirerProvider AcquirerProvider constructor
func NewAcquirerProvider(acquirers AcquirersStrategy) *AcquirerProvider {
	return &AcquirerProvider{acquirers}
}

// GetAcquirerStrategy returns a acquirer strategy
func (ap *AcquirerProvider) GetAcquirerStrategy(a Acquirer) IAcquirerStrategy {
	return ap.Acquirers[a]
}
