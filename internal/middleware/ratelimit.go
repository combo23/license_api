package middleware

import (
	"fmt"
	"os"
	"strconv"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(429, gin.H{"error": "Too many requests. Try again in " + time.Until(info.ResetTime).String()})
}

func mustConvertStringToInt(s string) uint {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string to int: %v", err))
	}
	return uint(i)
}

func NewRateLimitInstance() gin.HandlerFunc {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: mustConvertStringToInt(os.Getenv("RATE_LIMIT")),
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
	return mw
}
