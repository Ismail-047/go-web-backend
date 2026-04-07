package routers

import (
	"net/http"

	"backend-raw-http/internals/handlers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/heloo-world", handlers.HelloHandler)

	return mux
}