// +build wireinject

package userservicemock

import "github.com/google/wire"

func Build() *UserService {
	panic(wire.Build(ProviderSet))
}
