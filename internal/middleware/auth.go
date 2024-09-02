package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

const XApiKey = "X-API-Key"

var apiKey = os.Getenv("API_KEY")

func Auth(c *gin.Context) {
	apikey := c.GetHeader(XApiKey)
	if apikey != apiKey {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
