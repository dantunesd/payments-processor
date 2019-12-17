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

// IsValid validates payment data
func (p *Payment) IsValid() error {
	v := validator.New()
	if err := v.Struct(p); err != nil {
		return NewInvalidContentError(err.Error())
	}
	return nil
}

// Customer represents a payment's customer
type Customer struct {
	Name string `json:"name" validate:"required"`
}

// Details represents a payment's Details
type Details struct {
	Card         Card     `json:"card" validate:"required"`
	Amount       int      `json:"amount" validate:"min=100,required"`
	PaymentType  string   `json:"payment_type" validate:"required"`
	Installments int      `json:"installments" validate:"min=1,required"`
	Itens        []string `json:"itens" validate:"gte=1,required"`
}

// Card represents a payment's Card
type Card struct {
	SourceID        string `json:"source_id" validate:"required"`
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

// Source represents a open source
type Source struct {
	SourceID   string `json:"source_id"`
	CardNumber string `json:"card_number"`
	CVV        string `json:"cvv"`
}
