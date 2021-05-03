package authhandler

import (
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new AuthHandler containing the given AuthService.
func Provide(authService domain.AuthService, cookieService domain.CookieService, authMiddleware *authmiddleware.AuthMiddleware) *AuthHandler {
	return &AuthHandler{
		AuthService:    &authService,
		CookieService:  &cookieService,
		AuthMiddleware: authMiddleware,
	}
}
