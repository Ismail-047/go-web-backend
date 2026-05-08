package main

import (
	"fmt"
	"go-backend-postgresql/internals/api/router"
	"go-backend-postgresql/internals/config"
	"go-backend-postgresql/internals/database"
	"log"
	"net/http"
	"time"
)

func main() {

	cfg := config.Load()

	db := database.New(cfg)
	defer db.Close()

	r := router.New(db.Pool)

	server := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server is listening on http://localhost:3000")
	
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}