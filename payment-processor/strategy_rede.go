package payment

import "fmt"

import "context"

// RedeStrategy .
type RedeStrategy struct{}

// NewRedeStrategy .
func NewRedeStrategy() RedeStrategy {
	return RedeStrategy{}
}

// Process .
func (c RedeStrategy) Process(ctx context.Context, p Payment, s Source) error {
	fmt.Println("Processing Rede")
	return nil
}
