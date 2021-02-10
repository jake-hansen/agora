//+build wireinject

package userhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/providers"
	"github.com/jake-hansen/agora/services/userservice"
)

func Build() (*UserHandler, func(), error) {
	panic(wire.Build(ProviderProductionSet,
		             userservice.ProviderProductionSet,
		             userrepo.ProviderProductionSet,
		             database.ProviderProductionSet,
		             providers.ProductionSet))
}

func BuildTest(userService domain.UserService) *UserHandler {
	panic(wire.Build(ProviderTestSet))
}