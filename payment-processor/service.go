package payment

import "context"

// IService interface for service
type IService interface {
	ProcessPayment(ctx context.Context, p Payment) error
}

// Service implements the payment process
type Service struct {
	r ISourceRepository
}

// NewService constructor for Service
func NewService(r ISourceRepository) *Service {
	return &Service{
		r: r,
	}
}

// ProcessPayment process a payment
func (s Service) ProcessPayment(ctx context.Context, p Payment) error {

	if vErr := p.IsValid(); vErr != nil {
		return vErr
	}

	_, gErr := s.r.GetByID(ctx, p.Details.Card.SourceID)
	if gErr != nil {
		return gErr
	}

	return nil
}
