package userrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides provides a UserRepository configured with the given Manager.
func Provide(manager *database.Manager) *UserRepository {
	return &UserRepository{DB: manager.DB}
}

var (
	// ProviderProductionSet provides a UserRepository for Production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.UserRepository), new(*UserRepository)))
)
