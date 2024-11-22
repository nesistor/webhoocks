package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(ValidateTokenMiddleware)

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/api/webhook/time", app.GetTime)

	mux.Post("/api/webhook/placeholder", app.PlaceholderWebhook)

	return mux
}
