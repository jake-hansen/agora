package authhandler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/services/jwtservice"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// AuthHandler is the handler that manages authentication for the API.
type AuthHandler struct {
	AuthService *domain.AuthService
	CookieService *domain.CookieService
	AuthMiddleware  *authmiddleware.AuthMiddleware
}

// Register creates two endpoints to handle login and logout functionality.
// / (POST) - 	Login
// / (DELETE) - Logout
func (a *AuthHandler) Register(parentGroup *gin.RouterGroup) error {
	authGroup := parentGroup.Group("auth")
	{
		authGroup.POST("", a.Login)
		authGroup.DELETE("", a.Logout)
		authGroup.POST("/refresh", a.Refresh)
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

	// Retrieve tokenSet
	tokenSet, err := (*a.AuthService).Authenticate(*adapter.AuthDTOToDomain(&credentials))

	if err != nil {
		apiError := api.NewAPIError(http.StatusUnauthorized, err, "the provided credentials could not be validated")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
	} else {
		refreshCookieMaxAge := tokenSet.Refresh.Expires.Sub(time.Now()).Seconds()

		(*a.CookieService).SetCookie(c, "refresh", tokenSet.Refresh.Value, int(refreshCookieMaxAge), "/", true)
		c.JSON(http.StatusOK, adapter.TokenDomainToDTO(&tokenSet.Auth))
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
		return
	}

	c.Status(http.StatusOK)
}

func (a *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh")
	if err != nil {
		apiError := api.NewAPIError(http.StatusBadRequest, err, "the refresh token cookie could not be found or parsed")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return
	}

	authToken, err := (*a.AuthMiddleware).ParseToken(c.Request)
	if err != nil {
		apiError := api.NewAPIError(http.StatusBadRequest, err, "the access token could not be found or parsed")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return
	}

	tokenSet := domain.TokenSet{
		Auth:    *authToken,
		Refresh: domain.Token{
			Value:   refreshToken,
		},
	}

	newTokenSet, err := (*a.AuthService).RefreshToken(tokenSet)
	if err != nil {
		var jwtValidErr *jwt.ValidationError
		if errors.As(err, &jwtValidErr) {
			if errors.As(err.(*jwt.ValidationError).Inner, &jwtservice.ErrInvalidTokenPair) {
				err = api.NewAPIError(http.StatusBadRequest, err, "invalid token pair")
			} else {
				err = api.NewAPIError(http.StatusBadRequest, err, "error validating token(s)")
			}
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	} else {
		refreshCookieMaxAge := newTokenSet.Refresh.Expires.Sub(time.Now()).Seconds()

		(*a.CookieService).SetCookie(c, "refresh", newTokenSet.Refresh.Value, int(refreshCookieMaxAge), "/", true)
		c.JSON(http.StatusOK, adapter.TokenDomainToDTO(&newTokenSet.Auth))
	}
}
