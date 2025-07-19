package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS returns a gin.HandlerFunc that adds CORS headers
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// Set CORS headers
		c.Header("Access-Control-Allow-Origin", getAllowedOrigin(origin))
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// getAllowedOrigin returns the appropriate origin for CORS
func getAllowedOrigin(requestOrigin string) string {
	// In development, allow localhost origins
	allowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
	}

	// Check if request origin is in allowed list
	for _, allowed := range allowedOrigins {
		if requestOrigin == allowed {
			return requestOrigin
		}
	}

	// In production, return specific domain
	// For now, return the first allowed origin as default
	if len(allowedOrigins) > 0 {
		return allowedOrigins[0]
	}

	return "*"
}