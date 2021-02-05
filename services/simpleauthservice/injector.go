// +build wireinject

package simpleauthservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/providers"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/mocks/userservicemock"
	"github.com/jake-hansen/agora/services/userservice"
)

func Build() (*SimpleAuthService, func(), error) {
	panic(wire.Build(ProviderProductionSet, jwtservice.ProviderProductionSet, userservice.ProviderProductionSet, userrepo.ProviderProductionSet, database.ProviderProductionSet, providers.ProductionSet))
}

func BuildTest(jwtCfg jwtservice.Config) (*SimpleAuthService, error) {
	panic(wire.Build(ProviderTestSet, jwtservice.ProviderMockSet, userservicemock.ProviderSet))
}