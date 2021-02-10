//+build wireinject

package authmiddleware

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func BuildTest(authService domain.AuthService, parseToken ParseTokenFunc) *AuthMiddleware {
	panic(wire.Build(ProviderTestSet))
}
