// +build wireinject

package authservicemock

import "github.com/google/wire"

func Build() *AuthService {
	panic(wire.Build(ProviderSet))
}
