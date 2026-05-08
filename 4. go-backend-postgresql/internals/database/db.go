package database

import (
	"context"
	"fmt"
	"go-backend-postgresql/internals/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Service represents the database connection pool.
type Service struct {
	Pool *pgxpool.Pool
}

// New initializes a new PostgreSQL connection pool.
func New(cfg *config.Config) *Service {

	config, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	// Production-ready pool settings
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	// Verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	return &Service{Pool: pool}
}

// Close shuts down the pool gracefully.
func (s *Service) Close() {
	s.Pool.Close()
}