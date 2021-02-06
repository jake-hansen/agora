package authmiddleware

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(authService domain.AuthService) *AuthMiddleware {
	return New(&authService)
}

var (
	ProviderProductionSet = wire.NewSet(Provide)
	ProviderTestSet = wire.NewSet(Provide)
)
