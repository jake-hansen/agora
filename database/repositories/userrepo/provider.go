package userrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
)

func Provide(manager *database.Manager) *UserRepository {
	return &UserRepository{DB: manager.DB}
}

var (
	ProviderProductionSet = wire.NewSet(Provide)
	ProviderTestSet       = wire.NewSet(Provide)
)
