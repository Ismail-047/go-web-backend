package router

import (
	"go-backend-postgres-sqlc/internals/api/handlers"
	"go-backend-postgres-sqlc/internals/database/repository"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dbpool *pgxpool.Pool) *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	r.Get("/", handlers.Welcome)

	repo := repository.New(dbpool)
	userHandler := handlers.NewUserHandler(repo)


	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/welcome", handlers.Welcome)

		r.Mount("/users", UserRoutes(userHandler))
	})

	return r
}