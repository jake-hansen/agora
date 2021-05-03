package healthservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a HealthService configured with the provided SchemaMigrationRepo.
func Provide(repository domain.SchemaMigrationRepo) *HealthService {
	return &HealthService{schemaRepo: &repository}
}

var (
	// ProviderProdcutionSet provides a HealthService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.HealthService), new(*HealthService)))
)
