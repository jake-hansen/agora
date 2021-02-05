// +build wireinject

package jwtservicemock

import "github.com/google/wire"

func Build() *Service {
	panic(wire.Build(ProviderSet))
}
