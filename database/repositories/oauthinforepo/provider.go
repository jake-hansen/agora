package oauthinforepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides an OAuthInfoRepo configured by the given Manager.
func Provide(manager *database.Manager) *OAuthInfoRepo {
	return &OAuthInfoRepo{DB: manager.DB}
}

var (
	// ProviderSet provides an OAuthInfoRepo for use in production.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.OAuthInfoRepository), new(*OAuthInfoRepo)))
)
