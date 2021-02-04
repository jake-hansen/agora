// +build wireinject

package userservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*UserService, func(), error) {
	panic(wire.Build(Provide, userrepo.ProviderProductionSet, database.ProviderProductionSet, providers.ProductionSet))
}
