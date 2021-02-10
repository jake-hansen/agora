// +build wireinject

package jwtservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*JWTServiceImpl, error) {
	panic(wire.Build(ProviderProductionSet, providers.ProductionSet))
}

func BuildTest(cfg Config) (*JWTServiceImpl, error) {
	panic(wire.Build(ProviderTestSet))
}
