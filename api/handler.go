package main

import (
	"encoding/json"
	"net/http"
	"payments-processor/payment-processor"

	"github.com/go-chi/chi"
)

func createServerHandler() http.Handler {
	router := chi.NewRouter()
	router.Post("/payment/cielo", paymentCieloHandler())
	return router
}

func paymentCieloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payment := &payment.Payment{}
		if dErr := json.NewDecoder(r.Body).Decode(payment); dErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{dErr.Error()})
			return
		}

		if vErr := payment.Validate(); vErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{vErr.Error()})
			return
		}

		responseWriter(w, http.StatusOK, payment)
	}
}
