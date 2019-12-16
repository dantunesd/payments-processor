package payment

// Payment represents a incoming payment from client
type Payment struct {
	Customer      Customer      `json:"customer" validate:"required"`
	Card          Card          `json:"card"`
	Sale          Sale          `json:"sale"`
	Establishment Establishment `json:"establishment"`
}

// Customer represents a payment's customer
type Customer struct {
	Name string `json:"name"`
}

// Card represents a payment's card
type Card struct {
	Token      string `json:"token"`
	Brand      string `json:"brand"`
	Expiration string `json:"expiration"`
}

// Sale represents a payment's sale
type Sale struct {
	Amount       int      `json:"amount"`
	Installments int      `json:"installments"`
	Itens        []string `json:"itens"`
}

// Establishment represents a payment's establishment
type Establishment struct {
	Identifier string `json:"identifier"`
	Address    string `json:"address"`
	PostalCode string `json:"postalCode"`
}

// Validate validates payment data
func (p *Payment) Validate() error {
	return nil
}
