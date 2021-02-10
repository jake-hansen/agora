//+build wireinject

package authhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/providers"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/userservice"
)

func Build() (*AuthHandler, func(), error) {
	panic(wire.Build(ProviderProductionSet,
		             simpleauthservice.ProviderProductionSet,
		             jwtservice.ProviderProductionSet,
		             userservice.ProviderProductionSet,
		             userrepo.ProviderProductionSet,
		             database.ProviderProductionSet,
		             providers.ProductionSet))
}

func BuildTest(authService domain.AuthService) *AuthHandler {
	panic(wire.Build(ProviderTestSet))
}