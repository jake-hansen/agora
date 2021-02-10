package authhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

// Provide provides a new AuthHandler containing the given AuthService.
func Provide(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: &authService}
}

var (
	// ProviderProductionSet provides an AuthHandler for production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*AuthHandler)))

	// ProviderTestSet provides an AuthHandler for testing.
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*AuthHandler)))
)
