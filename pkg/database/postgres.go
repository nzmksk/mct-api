package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// NewPostgresConnection creates a new PostgreSQL database connection
func NewPostgresConnection() (*sql.DB, error) {
	// Build connection string from environment variables
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "postgres")
	password := getEnvOrDefault("DB_PASSWORD", "password")
	dbname := getEnvOrDefault("DB_NAME", "mct_api")
	sslmode := getEnvOrDefault("DB_SSLMODE", "disable")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)                 // Maximum connections
	db.SetMaxIdleConns(5)                  // Idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime
	db.SetConnMaxIdleTime(1 * time.Minute) // Idle timeout

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Successfully connected to PostgreSQL")
	return db, nil
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}