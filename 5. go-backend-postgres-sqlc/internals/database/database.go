package database

import (
	"context"
	"fmt"
	"go-backend-postgres-sqlc/internals/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) *Service {
	
	config, err := pgxpool.ParseConfig(cfg.DatabaseUrl);
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute
	
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	return &Service{Pool: pool}
}

func (s *Service) Close() {
	s.Pool.Close()
}