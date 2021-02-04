//+build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
)

func BuildSimpleAuthService() (*SimpleAuthService, error)  {
	panic(wire.Build(SimpleAuthServiceSet, UserServiceSet, userrepo.UserRepositorySet, config.ProvideViper))
}
