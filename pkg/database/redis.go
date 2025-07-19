package database

import (
	"context"
	"fmt"
	// "os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

// NewRedisConnection creates a new Redis client connection
func NewRedisConnection() (*redis.Client, error) {
	// Build Redis options from environment variables
	addr := getEnvOrDefault("REDIS_ADDR", "localhost:6379")
	password := getEnvOrDefault("REDIS_PASSWORD", "")
	
	dbStr := getEnvOrDefault("REDIS_DB", "0")
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB value: %w", err)
	}

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     10,
		MinIdleConns: 3,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	logrus.Info("Successfully connected to Redis")
	return rdb, nil
}
