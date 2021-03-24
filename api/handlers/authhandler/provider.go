package authhandler

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/cookieservice"
)

// Provide provides a new AuthHandler containing the given AuthService.
func Provide(authService domain.AuthService, cookieService *cookieservice.CookieService) *AuthHandler {
	return &AuthHandler{
		AuthService: &authService,
		CookieService: cookieService,
	}
}
