package payment

import "context"

// IService interface for service.
type IService interface {
	ProcessPayment(ctx context.Context, p Payment, acquirer Acquirer) error
}

// Service implements the payment process.
type Service struct {
	sr ISourceRepository
	ap IAcquirerProvider
}

// NewService Service's constructor.
func NewService(sr ISourceRepository, ap IAcquirerProvider) *Service {
	return &Service{
		sr: sr,
		ap: ap,
	}
}

// ProcessPayment process a payment with an acquirer.
func (s *Service) ProcessPayment(ctx context.Context, p Payment, acquirer Acquirer) error {

	if vErr := p.IsValid(); vErr != nil {
		return vErr
	}

	src, gErr := s.sr.GetByID(ctx, p.Details.Card.SourceID)
	if gErr != nil {
		return gErr
	}

	if pErr := s.ap.GetAcquirerStrategy(acquirer).Process(ctx, p, src); pErr != nil {
		return pErr
	}

	return nil
}
