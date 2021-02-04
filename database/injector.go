// +build wireinject

package database

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*Manager, func(), error) {
	panic(wire.Build(ProviderProductionSet, providers.ProductionSet))
}

func BuildTest(cfg Config) (*Manager, func(), error) {
	panic(wire.Build(ProviderTestSet))
}
