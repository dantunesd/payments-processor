package main

import (
	"encoding/json"
	"net/http"
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
