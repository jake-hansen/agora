package router

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/providers"
)

func Build() *Router {
	panic(wire.Build(ProviderProductionSet, providers.ProductionSet))
}

func BuildTest(cfg Config) *Router {
	panic(wire.Build(ProviderTestSet))
}
