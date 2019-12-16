package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func createServerHandler() http.Handler {
	router := chi.NewRouter()
	return router
}
