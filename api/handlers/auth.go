package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/domain"
	"net/http"
)

type AuthHandler struct {
	AuthService domain.AuthService
}

func NewAuthHandler(parentGroup *gin.RouterGroup, service domain.AuthService) {
	handler := &AuthHandler{AuthService: service}

	authGroup := parentGroup.Group("auth")
	{
		authGroup.POST("", handler.Login)
		authGroup.DELETE("", handler.Logout)
	}
}

func getCredentials(c *gin.Context) domain.Auth {
	var credentials domain.Auth
	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}
	return credentials
}

func (a *AuthHandler) Login(c *gin.Context) {
	credentials := getCredentials(c)

	// Retrieve token
	token, err := a.AuthService.Authenticate(credentials)

	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	c.JSON(http.StatusOK, token)
}

func (a *AuthHandler) Logout(c *gin.Context) {
	credentials := getCredentials(c)

	err := a.AuthService.Deauthenticate(credentials)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
	}

	c.Status(http.StatusOK)
}
