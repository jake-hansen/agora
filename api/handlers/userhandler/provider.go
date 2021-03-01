package userhandler

import (
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new UserHandler containing the given UserService.
func Provide(userService domain.UserService) *UserHandler {
	return &UserHandler{UserService: &userService}
}
