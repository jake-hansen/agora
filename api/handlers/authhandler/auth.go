package authhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/services/simpleauthservice"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

// AuthHandler is the handler that manages authentication for the API.
type AuthHandler struct {
	AuthService    *domain.AuthService
	CookieService  *domain.CookieService
	AuthMiddleware *authmiddleware.AuthMiddleware
}

// Register creates three endpoints to handle login, logout, and refresh functionality.
// / 		(POST)   - Login
// / 		(DELETE) - Logout
// /refresh (POST)   - Refresh
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
		refreshCookieMaxAge := tokenSet.Refresh.ExpiresAt.Sub(time.Now()).Seconds()

		(*a.CookieService).SetCookie(c, "refresh", string(tokenSet.Refresh.Value), int(refreshCookieMaxAge), "/", true)
		c.JSON(http.StatusOK, adapter.TokenDomainToDTO(&tokenSet.Auth))
	}
}

// Logout attempts to unauthenticate the given credentials retrieved from the body as JSON.
func (a *AuthHandler) Logout(c *gin.Context) {
	refreshTokenCookie, err := c.Cookie("refresh")
	if err != nil {
		apiError := api.NewAPIError(http.StatusBadRequest, err, "the refresh token cookie could not be found or parsed")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return
	}

	refreshToken := domain.TokenValue(refreshTokenCookie)

	err = (*a.AuthService).Deauthenticate(refreshToken)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}

// Refresh attempts to preform a refresh for a user's auth token.
func (a *AuthHandler) Refresh(c *gin.Context) {
	refreshTokenCookie, err := c.Cookie("refresh")
	if err != nil {
		apiError := api.NewAPIError(http.StatusBadRequest, err, "the refresh token cookie could not be found or parsed")
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return
	}

	refreshToken := domain.TokenValue(refreshTokenCookie)

	newTokenSet, err := (*a.AuthService).RefreshToken(refreshToken)
	if err != nil {
		var jwtValidErr *jwt.ValidationError
		var tokenReuseErr simpleauthservice.RefreshTokenReuse
		var tokenRevokedErr simpleauthservice.RefreshTokenRevoked
		if errors.As(err, &jwtValidErr) {
			err = api.NewAPIError(http.StatusBadRequest, err, "error validating token(s)")
		} else if errors.As(err, &tokenReuseErr) {
			err = api.NewAPIError(http.StatusForbidden, err, "refresh token reuse detected")
		} else if errors.As(err, &tokenRevokedErr) {
			err = api.NewAPIError(http.StatusForbidden, err, "refresh token was revoked")
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	} else {
		refreshCookieMaxAge := newTokenSet.Refresh.ExpiresAt.Sub(time.Now()).Seconds()

		(*a.CookieService).SetCookie(c, "refresh", string(newTokenSet.Refresh.Value), int(refreshCookieMaxAge), "/", true)
		c.JSON(http.StatusOK, adapter.TokenDomainToDTO(&newTokenSet.Auth))
	}
}
