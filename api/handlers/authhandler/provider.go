package authhandler

import (
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new AuthHandler containing the given AuthService.
func Provide(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: &authService}
}
