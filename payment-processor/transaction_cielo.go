package payment

import (
	"encoding/json"
	"net/http"
)

// Payment succeeded
const (
	Authorized       = 1
	PaymentConfirmed = 2
)

// CieloTransaction represents a cielo transaction
type CieloTransaction struct {
	r IResponser
}

// NewCieloTransaction constructor
func NewCieloTransaction(r IResponser) ITransaction {
	return CieloTransaction{
		r: r,
	}
}

// PaymentSucceeded returns the error
func (c CieloTransaction) PaymentSucceeded() error {
	if c.hasComunicationError() {
		return NewInternalServerError(string(c.r.GetBody()))
	}

	if c.hasIntegrationError() {
		c := c.decodeError()
		return NewTransactionError(c.Message)
	}

	out := c.decode()
	if c.hasEmissorError(out.Payment.Status) {
		return NewTransactionError(out.Payment.ReturnMessage)
	}

	return nil
}

func (c CieloTransaction) hasComunicationError() bool {
	return c.r.GetStatusCode() == http.StatusInternalServerError
}

func (c CieloTransaction) hasIntegrationError() bool {
	return c.r.GetStatusCode() == http.StatusBadRequest
}

func (c CieloTransaction) hasEmissorError(status int) bool {
	return !c.paymentSucceeded(status)
}

func (c CieloTransaction) paymentSucceeded(status int) bool {
	return status == PaymentConfirmed || status == Authorized
}

func (c CieloTransaction) decode() *CieloResponseBody {
	out := &CieloResponseBody{}
	json.Unmarshal(c.r.GetBody(), out)
	return out
}

func (c CieloTransaction) decodeError() *CieloIntegrationError {
	out := SCieloIntegrationError{}
	json.Unmarshal(c.r.GetBody(), &out)
	return out[0]
}
