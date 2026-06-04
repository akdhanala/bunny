package handlers

import (
    "errors"
    "net/http"

    "github.com/akdhanala/bunny/mapper"
    "github.com/akdhanala/bunny/entity"
    "github.com/akdhanala/bunny/controllers"
    "github.com/go-chi/chi/v5"

)


type bunnyHandler struct{
    controller *controllers.BunnyController
}

func NewBunnyHandler(cfg *entity.BunnyConfig) *bunnyHandler {
    controller := controllers.NewBunnyController(cfg)
    return &bunnyHandler{
        controller: controller,
    }
}

func (h *bunnyHandler) RegisterRoutes(r *chi.Mux) {
    r.Get("/search", h.handleRedirectRequest)
}

func (h *bunnyHandler) handleRedirectRequest(
    w http.ResponseWriter, 
    r *http.Request,
) {
    mappedRequest := mapper.HttpRequestToResolveDestinationRequest(r)
    destination, err := h.controller.ResolveDestination(mappedRequest)

    if (err != nil) {
        if unrecognizedError, ok := errors.AsType[*entity.CommandNotFound](err); ok {
            w.WriteHeader(unrecognizedError.Code())
        } else {
            w.WriteHeader(500)
        }

        w.Write([]byte(err.Error()))
        return
    }

    http.Redirect(w, r, destination, http.StatusFound)
}