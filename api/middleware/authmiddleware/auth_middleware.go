package authmiddleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/domain"
)

// ParseTokenFunc defines a function which parses an HTTP request for
// an authorization token.
type ParseTokenFunc = func(r *http.Request) (*domain.TokenValue, error)

// AuthMiddleware handles authentication by using AuthService to determine if
// a request is authenticated.
type AuthMiddleware struct {
	AuthService *domain.AuthService
	ParseToken  ParseTokenFunc
}

// New returns a new AuthMiddleware configured with the specified AuthService and ParseTokenFunc.
func New(authService *domain.AuthService, parseTokenFunc ParseTokenFunc) *AuthMiddleware {
	return &AuthMiddleware{
		AuthService: authService,
		ParseToken:  parseTokenFunc,
	}
}

// HandleAuth checks a request for valid authentication using the token obtained
// from the ParseToken function against the AuthService.
func (a *AuthMiddleware) HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := a.ParseToken(c.Request)

		if err != nil {
			apiError := api.NewAPIError(http.StatusUnauthorized, err, err.Error())
			_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
			c.Abort()
			return
		}

		_, err = (*a.AuthService).IsAuthenticated(*token)

		if err != nil {
			apiError := api.NewAPIError(http.StatusUnauthorized, err, "the provided token is not valid")
			_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUser attempts to get a User object from the provided Context. This
// is performed by passing the found AuthToken within the request (if any) to
// an AuthService which performs the User lookup.
func (a *AuthMiddleware) GetUser(c *gin.Context) (*domain.User, error) {
	token, err := a.ParseToken(c.Request)
	if err != nil {
		return nil, err
	}
	user, err := (*a.AuthService).GetUserFromAuthToken(*token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// getTokenFromBearerHeader parses the Authorization header for a Bearer token
// and returns the parsed result as a domain.AuthToken.
func getTokenFromBearerHeader(r *http.Request) (*domain.TokenValue, error) {
	t := r.Header.Get("Authorization")
	if t == "" {
		return nil, errors.New("token not found")
	}
	splitToken := strings.Split(t, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("could not parse token")
	}
	token := domain.TokenValue(splitToken[1])
	return &token, nil
}
