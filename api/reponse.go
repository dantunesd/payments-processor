package main

import (
	"encoding/json"
	"net/http"
	"payments-processor/payment-processor"
)

// ErrorResponse represents a response with error
type ErrorResponse struct {
	ErrorMessage string `json:"error"`
	Details      string `json:"details"`
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
		return getCodeByError(terr)
	default:
		return http.StatusInternalServerError
	}
}

func getCodeByError(err *payment.Error) int {
	switch err.ErrorType {
	case payment.InvalidContent:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
