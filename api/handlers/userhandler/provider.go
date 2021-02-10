package userhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

// Provide provides a new UserHandler containing the given UserService.
func Provide(userService domain.UserService) *UserHandler {
	return &UserHandler{UserService: &userService}
}

var (
	// ProviderProductionSet provides a UserHandler for production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*UserHandler)))

	// ProviderTestSet provides a UserHandler for testing.
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*UserHandler)))
)
