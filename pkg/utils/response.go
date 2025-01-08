package utils

import (
	"shopping-api-backend-go/internal/models"

	"github.com/gin-gonic/gin"
)

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, models.ErrorResponse{Error: message})
}

// RespondWithSuccess sends a standardized success response with data
func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}
