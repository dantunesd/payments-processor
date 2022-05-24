package payment

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Cielo Payment succeeded.
const (
	CieloPaymentAuthorized = 1
	CieloPaymentConfirmed  = 2
)

// CieloTransaction implementation of ITransaction.
type CieloTransaction struct {
	r IResponser
}

// NewCieloTransaction CieloTransaction's constructor.
func NewCieloTransaction(r IResponser) *CieloTransaction {
	return &CieloTransaction{
		r: r,
	}
}

// PaymentSucceeded verify the transaction results.
func (c *CieloTransaction) PaymentSucceeded() error {
	if c.hasComunicationError() {
		return NewInternalServerError(string(c.r.GetBody()))
	}

	if c.hasIntegrationError() {
		cir := c.decodeError()
		return NewTransactionError(fmt.Sprintf("ReturnCode: %d, ReturnMessage: %s", cir.Code, cir.Message))
	}

	crb := c.decode()
	if c.hasEmissorError(crb.Payment.Status) {
		return NewTransactionError(fmt.Sprintf("ReturnCode: %s, ReturnMessage: %s", crb.Payment.ReturnCode, crb.Payment.ReturnMessage))
	}

	return nil
}

func (c *CieloTransaction) hasComunicationError() bool {
	return c.r.GetStatusCode() == http.StatusInternalServerError
}

func (c *CieloTransaction) hasIntegrationError() bool {
	return c.r.GetStatusCode() == http.StatusBadRequest
}

func (c *CieloTransaction) hasEmissorError(status int) bool {
	return !c.paymentSucceeded(status)
}

func (c *CieloTransaction) paymentSucceeded(status int) bool {
	return status == CieloPaymentConfirmed || status == CieloPaymentAuthorized
}

func (c *CieloTransaction) decode() CieloResponseBody {
	crb := &CieloResponseBody{}
	json.Unmarshal(c.r.GetBody(), crb)
	return *crb
}

func (c *CieloTransaction) decodeError() CieloIntegrationError {
	cir := SCieloIntegrationError{}
	json.Unmarshal(c.r.GetBody(), &cir)
	if len(cir) > 0 {
		return *cir[0]
	}
	return CieloIntegrationError{}
}
