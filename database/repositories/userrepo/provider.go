package userrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

func Provide(manager *database.Manager) *UserRepository {
	return &UserRepository{DB: manager.DB}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.UserRepository), new(*UserRepository)))
	ProviderTestSet       = wire.NewSet(Provide, wire.Bind(new(domain.UserRepository), new(*UserRepository)))
)
