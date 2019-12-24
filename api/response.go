package main

import (
	"encoding/json"
	"net/http"
	"payments-processor/payment-processor"
)

// ErrorResponse represents a response with error
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func responseWriter(w http.ResponseWriter, code int, content interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	if content != nil {
		body, _ := json.Marshal(content)
		w.Write(body)
	}
}

func getHTTPCode(err error) int {
	switch terr := err.(type) {
	case *payment.Error:
		return terr.StatusCode
	default:
		return http.StatusInternalServerError
	}
}
