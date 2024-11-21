package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))

	// Time API route
	mux.Get("/api/webhook/time", app.GetTime)

	// Placeholder route
	mux.Post("/api/webhook/placeholder", app.PlaceholderWebhook)

	return mux
}
