package main

import (
	// "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"mct-api/internal/shared"
	"mct-api/pkg/database"
	"mct-api/pkg/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found")
	}

	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	
	// Set log level based on environment
	if os.Getenv("ENV") == "development" {
		logger.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		logger.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database connections
	db, err := database.NewPostgresConnection()
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to PostgreSQL")
	}
	defer db.Close()

	redis, err := database.NewRedisConnection()
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to Redis")
	}
	defer redis.Close()

	// Initialize Gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger(logger))
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, shared.APIResponse{
			Success: true,
			Data: map[string]string{
				"status": "ok",
				"service": "mct-api",
			},
		})
	})

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Authentication routes will be added here
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, shared.APIResponse{
				Success: true,
				Data:    "pong",
			})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.WithField("port", port).Info("Starting server")
	if err := router.Run(":" + port); err != nil {
		logger.WithError(err).Fatal("Failed to start server")
	}
}
