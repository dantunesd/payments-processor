package payment

import (
	"github.com/go-playground/validator/v10"
)

// Payment represents a incoming payment from client
type Payment struct {
	Customer      Customer      `json:"customer" validate:"required"`
	Card          Card          `json:"card" validate:"required"`
	Sale          Sale          `json:"sale" validate:"required"`
	Establishment Establishment `json:"establishment" validate:"required"`
}

// Customer represents a payment's customer
type Customer struct {
	Name string `json:"name" validate:"required"`
}

// Card represents a payment's card
type Card struct {
	Token      string `json:"token" validate:"required"`
	Brand      string `json:"brand" validate:"required"`
	Expiration string `json:"expiration" validate:"required"`
}

// Sale represents a payment's sale
type Sale struct {
	Amount       int      `json:"amount" validate:"required"`
	Installments int      `json:"installments" validate:"required"`
	Itens        []string `json:"itens" validate:"required"`
}

// Establishment represents a payment's establishment
type Establishment struct {
	Identifier string `json:"identifier" validate:"required"`
	Address    string `json:"address" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// IsValid validates payment data
func (p *Payment) IsValid() error {
	v := validator.New()
	return v.Struct(p)
}
