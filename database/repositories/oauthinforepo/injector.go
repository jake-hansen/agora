// +build wireinject

package oauthinforepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
)

func Build(manager *database.Manager) (*OAuthInfoRepo, func(), error) {
	panic(wire.Build(Provide))
}