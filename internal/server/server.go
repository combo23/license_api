package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"

	"license-api/internal/database"
	"license-api/internal/logger"
)

type Server struct {
	port   int
	logger zerolog.Logger
	db     database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	logger, _ := logger.NewLogger(fmt.Sprintf("logs/server-%v.log", time.Now()))

	NewServer := &Server{
		port:   port,
		logger: logger,
		db:     database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
