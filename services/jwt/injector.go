// +build wireinject

package jwt

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*Service, error) {
	panic(wire.Build(ProviderProductionSet, providers.ProductionSet))
}

func BuildTest() (*Service, error) {
	panic(wire.Build(ProviderTestSet))
}
