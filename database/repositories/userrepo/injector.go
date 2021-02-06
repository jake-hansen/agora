// +build wireinject

package userrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*UserRepository, func(), error) {
	panic(wire.Build(ProviderProductionSet, database.ProviderProductionSet, providers.ProductionSet))
}