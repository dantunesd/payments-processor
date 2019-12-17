package payment

// IService interface for service
type IService interface {
	ProcessPayment(p Payment) error
}

// Service implements the payment proccess
type Service struct {
	r ISourceRepository
}

// NewService constructor for Service
func NewService(r ISourceRepository) *Service {
	return &Service{
		r: r,
	}
}

// ProcessPayment proccess a payment
func (s Service) ProcessPayment(p Payment) error {
	return nil
}
