package router

import (
	"go-backend-postgres-sqlc/internals/api/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(h *handlers.UserHandler) * chi.Mux {

	r := chi.NewRouter()

	r.Post("/create-user", h.CreateUser)


	return r
}