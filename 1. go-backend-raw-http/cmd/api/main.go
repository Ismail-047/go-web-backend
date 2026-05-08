package main

import (
	"go-backend-raw-http/internals/api/router"
	"log"
	"net/http"
	"time"
)

func main() {

	mux := router.New()

	server := &http.Server{
		Addr: ":3000",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
	}

	log.Println("Server is listening on port 3000....")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}

}