package payment

import (
	"github.com/go-playground/validator/v10"
)

// Payment represents a incoming payment from client
type Payment struct {
	Customer      Customer      `json:"customer" validate:"required"`
	Details       Details       `json:"details" validate:"required"`
	Establishment Establishment `json:"establishment" validate:"required"`
}

// Customer represents a payment's customer
type Customer struct {
	Name string `json:"name" validate:"required"`
}

// Details represents a payment's Details
type Details struct {
	Source       Source   `json:"source" validate:"required"`
	Amount       int      `json:"amount" validate:"min=100,required"`
	PaymentType  string   `json:"payment_type" validate:"required"`
	Installments int      `json:"installments" validate:"min=1,required"`
	Itens        []string `json:"itens" validate:"gte=1,required"`
}

// Source represents a payment's Source
type Source struct {
	ID              string `json:"source_id" validate:"required"`
	Brand           string `json:"brand" validate:"required"`
	ExpirationYear  int    `json:"expiration_year" validate:"required"`
	ExpirationMonth int    `json:"expiration_month" validate:"required"`
}

// Establishment represents a payment's establishment
type Establishment struct {
	Identifier string `json:"identifier" validate:"required"`
	Address    string `json:"address" validate:"required"`
	PostalCode int    `json:"postal_code" validate:"required"`
}

// IsValid validates payment data
func (p *Payment) IsValid() error {
	v := validator.New()
	return v.Struct(p)
}
