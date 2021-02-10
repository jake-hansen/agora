package api

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// APIError represents an error that occurred during an operation on an endpoint.
type APIError struct {
	Status  int    // HTTP status returned to client.
	Err     error  // Error that occurred.
	Message string // Additional information about error returned to client.

}

// NewAPIError creates a new APIError with the given status, error, and message.
func NewAPIError(status int, err error, message string) *APIError {
	return &APIError{
		Status:  status,
		Err:     err,
		Message: message,
	}
}

// Error returns a string representation of the APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("%s (HTTP %d): %s", e.Message, e.Status, e.Err.Error())
}

// Unwrap returns the underlying error that caused the APIError.
func (e *APIError) Unwrap() error {
	return e.Err
}

// RegisterCustomValidation registers a custom tag name func to gin's validator
func RegisterCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// ValidationError represents an error that was caused by an invalid request.
type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

// DescriptiveValidationErrors returns a slice of ValidationErrors formatted in a descriptive way based on the given
// validator.ValidationErrors.
func DescriptiveValidationErrors(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
	}

	return errs
}

// SimpleValidationErrors returns a string map formatted in a descriptive way based on the given
// validator.ValidationErrors.
func SimpleValidationErrors(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
