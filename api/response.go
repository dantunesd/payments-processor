package main

import (
	"encoding/json"
	"net/http"
	"payments-processor/payment-processor"
)

// ErrorResponse represents a http response with error.
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// Response represents a http response with success.
type Response struct {
	Message string `json:"message"`
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
