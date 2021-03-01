package healthservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(repository domain.SchemaMigrationRepo) *HealthService {
	return &HealthService{schemaRepo: &repository}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.HealthService), new(*HealthService)))
)
