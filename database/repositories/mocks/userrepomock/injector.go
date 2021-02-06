// +build wireinject

package userrepomock

import "github.com/google/wire"

func Build() *UserRepository {
	panic(wire.Build(ProviderSet))
}
