package refreshtokenrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a RefreshTokenRepo configured with the provided Manager.
func Provide(manager *database.Manager) *RefreshTokenRepo {
	return &RefreshTokenRepo{DB: manager.DB}
}

var (
	// ProviderSet provides an RefreshTokenRepo for use in production.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.RefreshTokenRepository), new(*RefreshTokenRepo)))
)
