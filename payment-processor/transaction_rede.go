package payment

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// RedePaymentConfirmed succeeded.
const RedePaymentConfirmed = "00"

// RedeTransaction implementation of ITransaction.
type RedeTransaction struct {
	r IResponser
}

// NewRedeTransaction RedeTransaction's constructor.
func NewRedeTransaction(r IResponser) *RedeTransaction {
	return &RedeTransaction{
		r: r,
	}
}

// PaymentSucceeded verify the transaction results.
func (r *RedeTransaction) PaymentSucceeded() error {
	if r.hasComunicationError() {
		return NewInternalServerError(string(r.r.GetBody()))
	}

	rrb := r.decode()
	if !r.paymentSucceeded(rrb.ReturnCode) {
		return NewTransactionError(fmt.Sprintf("ReturnCode: %s, ReturnMessage: %s", rrb.ReturnCode, rrb.ReturnMessage))
	}

	return nil
}

func (r *RedeTransaction) hasComunicationError() bool {
	return r.r.GetStatusCode() == http.StatusInternalServerError
}

func (r *RedeTransaction) paymentSucceeded(code string) bool {
	return code == RedePaymentConfirmed
}

func (r *RedeTransaction) decode() RedeResponseBody {
	rrb := &RedeResponseBody{}
	json.Unmarshal(r.r.GetBody(), rrb)
	return *rrb
}
