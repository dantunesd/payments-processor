package main

import (
	"encoding/json"
	"net/http"
	"payments-processor/payment-processor"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createServerHandler(s payment.IService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Post("/payment/cielo", paymentCieloHandler(s))
	return r
}

func paymentCieloHandler(s payment.IService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payment := &payment.Payment{}
		if dErr := json.NewDecoder(r.Body).Decode(payment); dErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{"Invalid body content", dErr.Error()})
			return
		}

		if vErr := payment.IsValid(); vErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{"There are invalid values or missing required fields. Please, see the documentation", vErr.Error()})
			return
		}

		if pErr := s.ProcessPayment(*payment); pErr != nil {
			responseWriter(w, http.StatusBadRequest, &ErrorResponse{"Failed to proccess payment", pErr.Error()})
			return
		}

		responseWriter(w, http.StatusOK, payment)
	}
}
