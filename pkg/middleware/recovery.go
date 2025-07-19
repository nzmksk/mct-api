package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"mct-api/internal/shared"
)

// Recovery returns a gin.HandlerFunc that recovers from panics
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// Log the panic with stack trace
		logrus.WithFields(logrus.Fields{
			"panic":      recovered,
			"stack":      string(debug.Stack()),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Error("Panic recovered")

		// Return standardized error response
		c.JSON(http.StatusInternalServerError, shared.APIResponse{
			Success: false,
			Error: &shared.APIError{
				Code:    string(shared.ErrInternalError),
				Message: "An unexpected error occurred",
			},
		})
	})
}