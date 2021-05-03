package schemamigrationrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a SchemaMigrationRepo configured with the provided Manager.
func Provide(manager *database.Manager) *SchemaMigrationRepo {
	return &SchemaMigrationRepo{DB: manager.DB}
}

var (
	// ProviderProductionSet provides a SchemaMigrationRepo for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.SchemaMigrationRepo), new(*SchemaMigrationRepo)))
)
