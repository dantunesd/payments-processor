package payment

import "fmt"

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
