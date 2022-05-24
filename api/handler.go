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
	r.Post("/payment/cielo", paymentHandler(s, payment.Cielo))
	r.Post("/payment/rede", paymentHandler(s, payment.Rede))
	return r
}

func paymentHandler(s payment.IService, acquirer payment.Acquirer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		p := &payment.Payment{}

		if dErr := json.NewDecoder(r.Body).Decode(p); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{"Invalid body content", dErr.Error()})
			return
		}

		if pErr := s.ProcessPayment(r.Context(), *p, acquirer); pErr != nil {
			responseWriter(w, getHTTPCode(pErr), ErrorResponse{"Failed to proccess payment", pErr.Error()})
			return
		}

		responseWriter(w, http.StatusOK, Response{"payment succeeded"})
	}
}
