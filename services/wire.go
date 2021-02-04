//+build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/userservice"
)

func BuildSimpleAuthService() (*simpleauthservice.SimpleAuthService, error)  {
	panic(wire.Build(simpleauthservice.SimpleAuthServiceSet, userservice.UserServiceSet, userrepo.UserRepositorySet, config.ProvideViper))
}
