package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "X-Request-Id"

// RequestID is a middleware that adds a unique request ID to the context for each request. Usefull for logging.
func RequestID(c *gin.Context) {
	requestID := c.Request.Header.Get(string(RequestIDKey))
	if requestID == "" {
		requestID = uuid.New().String()
	}
	c.Set(string(RequestIDKey), requestID)
	c.Writer.Header().Set("X-Request-Id", requestID)
	c.Next()
}

func GetRequestID(c *gin.Context) string {
	return c.Writer.Header().Get(RequestIDKey)
}
