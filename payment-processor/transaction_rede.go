package payment

import (
	"encoding/json"
	"net/http"
)

// RedeTransaction represents a R transaction
type RedeTransaction struct {
	r IResponser
}

// NewRedeTransaction constructor
func NewRedeTransaction(r IResponser) RedeTransaction {
	return RedeTransaction{
		r: r,
	}
}

// IsSucceeded validates if the transaction was succeded
func (r RedeTransaction) IsSucceeded() bool {
	if r.hasComunicationError() {
		return false
	}

	out := r.decode()

	return r.paymentSucceeded(out.ReturnCode)
}

// GetError returns the error
func (r RedeTransaction) GetError() error {
	if r.hasComunicationError() {
		return NewInternalServerError(string(r.r.GetBody()))
	}

	out := r.decode()
	if !r.paymentSucceeded(out.ReturnCode) {
		return NewTransactionError(out.ReturnMessage)
	}

	return nil
}

func (r RedeTransaction) hasComunicationError() bool {
	return r.r.GetStatusCode() == http.StatusInternalServerError
}

func (r RedeTransaction) paymentSucceeded(code string) bool {
	return code == RedePaymentConfirmed
}

func (r RedeTransaction) decode() *RedeResponseBody {
	out := &RedeResponseBody{}
	json.Unmarshal(r.r.GetBody(), out)
	return out
}
