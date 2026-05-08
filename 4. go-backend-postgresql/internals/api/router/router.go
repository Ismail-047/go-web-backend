package router

import (
	"go-backend-postgresql/internals/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbpool *pgxpool.Pool) *chi.Mux {

	r := chi.NewRouter()


	// Production Middleware Stack
	r.Use(middleware.RequestID) // Injects a request ID into the context
	r.Use(middleware.RealIP)    // Trust reverse proxy headers (crucial for PaaS)
	r.Use(middleware.Logger)    // Log API requests
	r.Use(middleware.Recoverer) // Recover from panics without crashing the server

	// Other Handlers
	userHanler := handlers.NewUserHandler(dbpool)


	// --- Root Route ---
	r.Get("/", handlers.Welcome)


	// Route Grouping (API Versioning)
	r.Route("/api/v1", func(r chi.Router) {
		
		// Chi uses explicit method functions instead of generic HandleFunc
		r.Get("/test/hello-world", handlers.HelloWorld)

		r.Post("/users", userHanler.Create)
	})

	return r
}