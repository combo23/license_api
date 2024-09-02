package server

import "github.com/gin-gonic/gin"

// Error Object to handle errors
type AppError struct {
	StatusCode int    // HTTP status code
	Message    string // Error message to return to the user
	LogMessage string // Message to log internally
}

var InternalServerError = func(log string) *AppError {
	return &AppError{
		StatusCode: 500,
		Message:    "Internal server error",
		LogMessage: log,
	}
}

var InvalidLicense = func(reason string) *AppError {
	return &AppError{
		StatusCode: 400,
		Message:    "Invalid license",
		LogMessage: "Invalid license: " + reason,
	}
}

var InvalidPayload = func(reason string) *AppError {
	return &AppError{
		StatusCode: 400,
		Message:    "Invalid payload: " + reason,
		LogMessage: "Invalid payload: " + reason,
	}
}

// HandleError handles the error and returns the error message to the user and logs it
func (s *Server) HandleError(c *gin.Context, err *AppError) {
	s.logger.Error().Msg(err.LogMessage)
	c.JSON(err.StatusCode, gin.H{"error": err.Message})
	c.Abort()
}
