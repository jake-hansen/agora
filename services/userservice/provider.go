package userservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(repository domain.UserRepository) *UserService {
	return &UserService{repo: repository}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))
	ProviderTestSet		  = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))
)
