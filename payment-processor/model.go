package payment

import (
	"github.com/go-playground/validator/v10"
)

// Payment represents a incoming payment from client.
type Payment struct {
	OrderID       string        `json:"order_id" validate:"required"`
	Customer      Customer      `json:"customer" validate:"required"`
	Details       Details       `json:"details" validate:"required"`
	Establishment Establishment `json:"establishment" validate:"required"`
}

// IsValid validates payment data.
func (p *Payment) IsValid() error {
	v := validator.New()
	if err := v.Struct(p); err != nil {
		return NewInvalidContentError(err.Error())
	}
	return nil
}

// Customer represents a payment's customer.
type Customer struct {
	Name string `json:"name" validate:"required"`
}

// Details represents a payment's details.
type Details struct {
	Card         Card     `json:"card" validate:"required"`
	Amount       int      `json:"amount" validate:"min=100,required"`
	PaymentType  string   `json:"payment_type" validate:"alpha,required"`
	Installments int      `json:"installments" validate:"min=1,required"`
	Itens        []string `json:"itens" validate:"gte=1,required"`
}

// Card represents a payment's card.
type Card struct {
	SourceID        string `json:"source_id" validate:"alphanum,required"`
	Brand           string `json:"brand" validate:"alpha,required"`
	ExpirationYear  int    `json:"expiration_year" validate:"required"`
	ExpirationMonth int    `json:"expiration_month" validate:"min=1,max=12,required"`
}

// Establishment represents a payment's establishment.
type Establishment struct {
	Identifier string `json:"identifier" validate:"required"`
	Address    string `json:"address" validate:"required"`
	PostalCode int    `json:"postal_code" validate:"required"`
}

// Source represents a open source.
type Source struct {
	SourceID   string `json:"source_id"`
	CardNumber string `json:"card_number"`
	CVV        int    `json:"cvv"`
}

// CieloRequestBody represents the request body of Cielo.
type CieloRequestBody struct {
	MerchantOrderID string              `json:"MerchantOrderId"`
	Customer        Customer            `json:"Customer"`
	Payment         CieloRequestPayment `json:"Payment"`
}

// CieloRequestPayment represents Cielo's payment node.
type CieloRequestPayment struct {
	Type         string          `json:"Type"`
	Amount       int             `json:"Amount"`
	Installments int             `json:"Installments"`
	CreditCard   CieloCreditCard `json:"CreditCard"`
}

// CieloCreditCard represents Cielo's creditcard node.
type CieloCreditCard struct {
	CardNumber     string `json:"CardNumber"`
	Holder         string `json:"Holder"`
	ExpirationDate string `json:"ExpirationDate"`
	SecurityCode   int    `json:"SecurityCode"`
	Brand          string `json:"Brand"`
}

// CieloResponseBody represents a Cielo's response body with success.
type CieloResponseBody struct {
	Payment CieloPaymentResponse `json:"Payment"`
}

// CieloPaymentResponse represents Cielo's payment node.
type CieloPaymentResponse struct {
	Status        int    `json:"Status"`
	ReturnMessage string `json:"ReturnMessage"`
	ReturnCode    string `json:"ReturnCode"`
}

// SCieloIntegrationError is a slice of CieloIntegrationError.
type SCieloIntegrationError []*CieloIntegrationError

// CieloIntegrationError represents a Cielo's response body with error.
type CieloIntegrationError struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

// RedeRequestBody represents a Rede's request body.
type RedeRequestBody struct {
	Capture         bool   `json:"capture"`
	Reference       string `json:"reference"`
	Amount          int    `json:"amount"`
	HolderName      string `json:"cardholderName"`
	CardNumber      string `json:"cardNumber"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  int    `json:"expirationYear"`
	SecurityCode    string `json:"securityCode"`
	Installments    int    `json:"installments"`
}

// RedeResponseBody represents a Rede's response body.
type RedeResponseBody struct {
	ReturnMessage string `json:"returnMessage"`
	ReturnCode    string `json:"returnCode"`
}
