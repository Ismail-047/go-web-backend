package router

import (
	"go-backend-raw-http/internals/api/handlers"
	"net/http"
)

func New() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.HelloWorld)
	mux.HandleFunc("GET /api/v1/test/hello-world", handlers.HelloWorld)

	return  mux
}