package userhandler

import (
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new UserHandler containing the given UserService and AuthMiddleware.
func Provide(userService domain.UserService, authMiddleware *authmiddleware.AuthMiddleware) *UserHandler {
	return &UserHandler{
		UserService:    &userService,
		AuthMiddleware: authMiddleware,
	}
}
