package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"backend-raw-http/internals/routers"
)

func main() {

	appRouter := routers.New()

	server := &http.Server{
		Addr: ":3000",
		Handler: appRouter,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	fmt.Println("Server is listening on port 3000....")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}