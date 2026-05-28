package handlers

import (
    "net/http"

    "github.com/akdhanala/bunny/controllers"
    "github.com/go-chi/chi/v5"
)


type bunnyHandler struct{
    controller *controllers.BunnyController
}

func NewBunnyHandler() *bunnyHandler {
    return &bunnyHandler{}
}

func (h *bunnyHandler) RegisterRoutes(r *chi.Mux) {
    r.Get("/bunny", h.handleRedirectRequest)
}

func (h *bunnyHandler) handleRedirectRequest(
    w http.ResponseWriter, 
    r *http.Request,
) {
    destination := h.controller.ResolveDestination()
    http.Redirect(w, r, destination, http.StatusFound)
}