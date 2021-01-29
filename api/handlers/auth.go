package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/domain"
	"net/http"
)

// AuthHandler is the handler that manages authentication for the API.
type AuthHandler struct {
	AuthService *domain.AuthService
}

// NewAuthHandler registers a new AuthHandler on the specified RouterGroup. The AuthHandler uses
// the specified AuthService.
func NewAuthHandler(parentGroup *gin.RouterGroup, service domain.AuthService) {
	handler := &AuthHandler{AuthService: &service}

	authGroup := parentGroup.Group("auth")
	{
		authGroup.POST("", handler.Login)
		authGroup.DELETE("", handler.Logout)
	}
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
	var credentials domain.Auth
	err := c.ShouldBind(&credentials)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Retrieve token
	token, err := (*a.AuthService).Authenticate(credentials)

	if err != nil {
		apiError := api.NewAPIError(http.StatusUnauthorized, err, "the provided credentials could not be validated")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
	} else {
		c.JSON(http.StatusOK, token)
	}
}

// Logout attempts to unauthenticate the given credentials retrieved from the body as JSON.
func (a *AuthHandler) Logout(c *gin.Context) {
	var token domain.Token
	err := c.ShouldBind(&token)
	err = validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	err = (*a.AuthService).Deauthenticate(token)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	c.Status(http.StatusOK)
}
