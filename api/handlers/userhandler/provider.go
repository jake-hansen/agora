package userhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

func Provide(userService domain.UserService) *UserHandler  {
	return &UserHandler{UserService: &userService}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*UserHandler)))
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*UserHandler)))
)
