package authmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/domain"
	"net/http"
	"strings"
)

// ParseTokenFunc defines a function which parses an HTTP request for
// an authorization token.
type ParseTokenFunc = func(r *http.Request) domain.Token

// AuthMiddleware handles authentication by using AuthService to determine if
// a request is authenticated.
type AuthMiddleware struct {
	AuthService		*domain.AuthService
	ParseToken		ParseTokenFunc
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
		token := a.ParseToken(c.Request)
		_, err := (*a.AuthService).IsAuthenticated(token)
		if err != nil {
			apiError := &api.APIError{
				Status:  401,
				Err:     err,
				Message: "the provided token is not valid",
			}
			_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		} else {
			c.Next()
		}
	}
}

// getTokenFromBearerHeader parses the Authorization header for a Bearer token
// and returns the parsed result as a domain.Token.
func getTokenFromBearerHeader(r*http.Request) domain.Token {
	t := r.Header.Get("Authorization")
	splitToken := strings.Split(t, "Bearer ")
	token := domain.Token{Value: splitToken[1]}
	return token
}