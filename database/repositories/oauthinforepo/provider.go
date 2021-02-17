package oauthinforepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

func Provide(manager *database.Manager) *OAuthInfoRepo {
	return &OAuthInfoRepo{DB: manager.DB}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.OAuthInfoRepository), new(*OAuthInfoRepo)))
)