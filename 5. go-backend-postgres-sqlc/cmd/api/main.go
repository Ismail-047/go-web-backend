/*
GO BACKEND WITH
  - POSTGRESQL
  - CHI ROUTER
  - SQLC - MIGRATIONS, QUERIES, REPOSITORY
  - SEPRATING ROUTES IN DIFF FILES
  - FRONTEND REQUEST VALIDATIONS
  - ENV VALIDATIONS
  - CUSTOM SUCCESS AND ERROR RESPONSES
  - DATABSE MIGRATIONS
  - HASHING PASSWORD
  - AIR (RESTART SERVER ON EVERY CHANGE - GOLAND NODEMON)
*/
package main

import (
	"fmt"
	"go-backend-postgres-sqlc/internals/api/router"
	"go-backend-postgres-sqlc/internals/config"
	"go-backend-postgres-sqlc/internals/database"
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
		log.Fatalf("Error listening server: %v", err)
	}
}