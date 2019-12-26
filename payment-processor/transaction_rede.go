package payment

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RedePaymentConfirmed succeeded
const RedePaymentConfirmed = "00"

// RedeTransaction represents a Rede transaction
type RedeTransaction struct {
	r IResponser
}

// NewRedeTransaction constructor
func NewRedeTransaction(r IResponser) RedeTransaction {
	return RedeTransaction{
		r: r,
	}
}

// PaymentSucceeded returns the error
func (r RedeTransaction) PaymentSucceeded() error {
	if r.hasComunicationError() {
		return NewInternalServerError(string(r.r.GetBody()))
	}

	out := r.decode()
	if !r.paymentSucceeded(out.ReturnCode) {
		return NewTransactionError(fmt.Sprintf("ReturnCode: %s, ReturnMessage: %s", out.ReturnCode, out.ReturnMessage))
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
