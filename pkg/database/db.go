package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(DATABASE_URL string) *pgxpool.Pool {
	// Create a new connection pool using pgx
	pool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatalf("💻 Failed to connect to database: %v", err)
	}

	// Verify connection
	if pingErr := pool.Ping(context.Background()); pingErr != nil {
		log.Fatalf("📞 Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully!")
	return pool
}
