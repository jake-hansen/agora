// +build wireinject

package log

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/providers"
)

func Build() (*Log, func(), error) {
	panic(wire.Build(ProviderProductionSet, providers.ProductionSet))
}
