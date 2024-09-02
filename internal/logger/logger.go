package logger

import (
	"license-api/internal/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func override(event *zerolog.Event, c *gin.Context) *zerolog.Event {
	return event.
		Str("path", c.FullPath()).
		Str("ip", c.ClientIP()).
		Str("request_id", middleware.GetRequestID(c))
}

func LogError(logger zerolog.Logger, c *gin.Context, msg string) {
	override(logger.Error(), c).Msg(msg)
}

func LogInfo(logger zerolog.Logger, c *gin.Context, msg string) {
	override(logger.Info(), c).Msg(msg)
}

func LogDebug(logger zerolog.Logger, c *gin.Context, msg string) {
	override(logger.Debug(), c).Msg(msg)
}

func LogWarn(logger zerolog.Logger, c *gin.Context, msg string) {
	override(logger.Warn(), c).Msg(msg)
}

func NewLogger(path string) (zerolog.Logger, error) {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	fileWriter, err := os.Create(path)
	if err != nil {
		return zerolog.Logger{}, err
	}
	jsonWriter := zerolog.New(fileWriter).With().Timestamp().Logger()

	multiWriter := zerolog.MultiLevelWriter(consoleWriter, jsonWriter)

	return zerolog.New(multiWriter).Level(zerolog.TraceLevel).With().Timestamp().Logger(), nil
}
