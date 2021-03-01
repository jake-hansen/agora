package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/api"
	"net/http"
)

// APIErrorJSON represents an error message.
type APIErrorJSON struct {
	Error string `json:"error"`
}

// PublicErrorHandler middleware handles public errors for the Gin framework.
func PublicErrorHandler() gin.HandlerFunc {
	return handlePublicErrors()
}

// handlePublicErrors reports errors to the client in a meaningful way.
// If a ValidationError is available, they will be reported to the client.
// If an APIError is available, the provided error message will be returned
// to the client along with the provided HTTP status. If an APIError is not
// available, a generic error message is returned along with a 500 status.
func handlePublicErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			var apiError *api.APIError
			var verr validator.ValidationErrors
			if errors.As(err.Err, &verr) {
				c.JSON(http.StatusBadRequest, gin.H{"validation errors": api.DescriptiveValidationErrors(verr)})
			} else if errors.As(err.Err, &apiError) {
				displayError := APIErrorJSON{
					Error: apiError.Message,
				}
				c.JSON(apiError.Status, displayError)
			} else {
				displayError := APIErrorJSON{
					Error: "unknown error occurred.",
				}
				c.JSON(http.StatusInternalServerError, displayError)
			}
		}
	}
}