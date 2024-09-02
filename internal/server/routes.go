package server

import (
	"license-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(middleware.RequestID)
	r.Use(middleware.NewRateLimitInstance())

	r.GET("/health", s.healthHandler)

	api := r.Group("/license")
	api.GET("/:id", middleware.Auth, s.getLicense)
	api.POST("/create", middleware.Auth, s.createLicense)
	api.POST("/verify", s.verifyLicense)
	api.POST("/unbind/:id", middleware.Auth, s.unbindLicense)
	api.POST("/ban/:id", middleware.Auth, s.banLicense)
	api.POST("/update", middleware.Auth, s.updateLicense)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
