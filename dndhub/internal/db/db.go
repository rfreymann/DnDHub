package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return db
}
