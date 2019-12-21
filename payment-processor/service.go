package payment

import "context"

// IService interface for service
type IService interface {
	ProcessPayment(ctx context.Context, p Payment, a Acquirer) error
}

// Service implements the payment process
type Service struct {
	r ISourceRepository
	a IAcquirerProvider
}

// NewService constructor for Service
func NewService(r ISourceRepository, a IAcquirerProvider) *Service {
	return &Service{
		r: r,
		a: a,
	}
}

// ProcessPayment process a payment
func (s Service) ProcessPayment(ctx context.Context, p Payment, ac Acquirer) error {

	if vErr := p.IsValid(); vErr != nil {
		return vErr
	}

	src, gErr := s.r.GetByID(ctx, p.Details.Card.SourceID)
	if gErr != nil {
		return gErr
	}

	if pErr := s.a.GetAcquirerStrategy(ac).Process(p, src); pErr != nil {
		return pErr
	}

	return nil
}
