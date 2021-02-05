package authhandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

// AuthHandler is the handler that manages authentication for the API.
type AuthHandler struct {
	AuthService *domain.AuthService
}

func (a *AuthHandler) Register(parentGroup *gin.RouterGroup) error {
	authGroup := parentGroup.Group("auth")
	{
		authGroup.POST("", a.Login)
		authGroup.DELETE("", a.Logout)
	}
	return nil
}

func validateHelper(err error) error {
	var verr validator.ValidationErrors
	if err != nil && !errors.As(err, &verr) {
		err = api.NewAPIError(http.StatusBadRequest, err, "could not parse request body")
	}
	return err
}

// Login attempts to authenticate the given credentials retrieved from the body as JSON.
func (a *AuthHandler) Login(c *gin.Context) {
	var credentials dto.Auth
	err := c.ShouldBind(&credentials)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Retrieve token
	token, err := (*a.AuthService).Authenticate(*adapter.AuthDTOToDomain(&credentials))

	if err != nil {
		apiError := api.NewAPIError(http.StatusUnauthorized, err, "the provided credentials could not be validated")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
	} else {
		c.JSON(http.StatusOK, adapter.TokenDomainToDTO(token))
	}
}

// Logout attempts to unauthenticate the given credentials retrieved from the body as JSON.
func (a *AuthHandler) Logout(c *gin.Context) {
	var token dto.Token
	err := c.ShouldBind(&token)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	err = (*a.AuthService).Deauthenticate(*adapter.TokenDTOToDomain(&token))
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	c.Status(http.StatusOK)
}