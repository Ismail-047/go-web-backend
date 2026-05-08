package main

import (
	"go-backend-chi/internals/api/router"
	"log"
	"net/http"
	"time"
)

func main() {

	r := router.New()

	server := &http.Server{
		Addr: ":3000",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}


	log.Println("Server is listenting on port 3000...")

	if err:= server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server Error: %v", err)
	}
}