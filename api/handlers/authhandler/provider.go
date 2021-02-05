package authhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

func Provide(authService domain.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: &authService}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*AuthHandler)))
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*AuthHandler)))
)
