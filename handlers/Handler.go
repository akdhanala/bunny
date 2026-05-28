package handlers

import "github.com/go-chi/chi/v5"

type Handler interface {
	RegisterRoutes(*chi.Mux)
}