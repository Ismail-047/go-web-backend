package router

import (
	"go-backend-chi/internals/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	// 1. Production Middleware Stack
	r.Use(middleware.RequestID) // Injects a request ID into the context
	r.Use(middleware.RealIP)    // Trust reverse proxy headers (crucial for PaaS)
	r.Use(middleware.Logger)    // Log API requests
	r.Use(middleware.Recoverer) // Recover from panics without crashing the server

	// --- Root Route ---
	r.Get("/", handlers.Welcome)

	// 2. Route Grouping (API Versioning)
	r.Route("/api/v1", func(r chi.Router) {
		// Chi uses explicit method functions instead of generic HandleFunc
		r.Get("/test/hello-world", handlers.HelloWorld)
	})

	return r
}