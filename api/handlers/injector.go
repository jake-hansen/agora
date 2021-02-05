// +build wireinject

package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/router/handlers"
)

func Build() ([]handlers.Handler, func(), error) {
	panic(wire.Build(ProviderProductionSet))
}
