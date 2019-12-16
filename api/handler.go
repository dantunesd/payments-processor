package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func createServerHandler() http.Handler {
	router := chi.NewRouter()
	router.Post("/payment/cielo", paymentCieloHandler())
	return router
}

func paymentCieloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
	}
}
