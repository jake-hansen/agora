//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/providers"
	"github.com/jake-hansen/agora/router"
)

func Build() (*Server, func(), error) {
	panic(wire.Build(ProviderProductionSet,
			         router.ProviderProductionSet,
			         handlers.ProviderProductionSet,
			         middleware.ProviderProductionSet,
			         providers.ProductionSet))
}
