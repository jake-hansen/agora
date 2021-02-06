package authmiddleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/domain"
	"strings"
)

type AuthMiddleware struct {
	AuthService		*domain.AuthService
}

func New(authService *domain.AuthService) *AuthMiddleware {
	return &AuthMiddleware{AuthService: authService}
}

func (a *AuthMiddleware) HandleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("Authorization")
		splitToken := strings.Split(t, "Bearer ")
		token := domain.Token{Value: splitToken[1]}
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