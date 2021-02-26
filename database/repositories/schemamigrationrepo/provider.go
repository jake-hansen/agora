package schemamigrationrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

func Provide(manager *database.Manager) *SchemaMigrationRepo {
	return &SchemaMigrationRepo{DB: manager.DB}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.SchemaMigrationRepo), new(*SchemaMigrationRepo)))
)
