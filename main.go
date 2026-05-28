package main

import (
    "os"
    "net/http"

    "github.com/akdhanala/bunny/entity"
	"github.com/akdhanala/bunny/handlers"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "gopkg.in/yaml.v3"
)

func main() {
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        os.Exit(1)
    }

    var rawCfg entity.BunnyConfig
    if err := yaml.Unmarshal(data, &rawCfg); err != nil { 
        os.Exit(1)
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

	registeredHandlers := []handlers.Handler{
        handlers.NewBunnyHandler(&rawCfg),
    }

    for _, handler := range registeredHandlers {
        handler.RegisterRoutes(r)
    }

    http.ListenAndServe(":8080", r)
}
