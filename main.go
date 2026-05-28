package main

import (
    "net/http"

	"github.com/akdhanala/bunny/handlers"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

	registeredHandlers := []handlers.Handler{
        handlers.NewBunnyHandler(),
    }

    for _, handler := range registeredHandlers {
        handler.RegisterRoutes(r)
    }

    http.ListenAndServe(":8080", r)
}
