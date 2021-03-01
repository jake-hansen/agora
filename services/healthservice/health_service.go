package healthservice

import (
	"errors"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

var neededSchemaVersion = 7

type HealthService struct {
	schemaRepo *domain.SchemaMigrationRepo
}

func (h *HealthService) GetHealth() (*domain.Health, error) {
	return h.backwardsCompatibleVersionHealthCheck(7)
}

func (h *HealthService) specificVersionNeededHealthCheck(neededVersion int) (*domain.Health, error) {
	health := &domain.Health{Healthy: false}

	_, err := (*h.schemaRepo).GetSchemaMigrationByVersion(neededVersion)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			health.Reason = fmt.Sprintf("database version is not at needed version %d", neededSchemaVersion)
		} else {
			return nil, err
		}
	} else {
		health.Healthy = true
	}

	return health, nil
}

func (h *HealthService) backwardsCompatibleVersionHealthCheck(neededVersion int) (*domain.Health, error) {
	health := &domain.Health{Healthy: false}

	schema, err := (*h.schemaRepo).GetSchemaMigration()
	if err != nil {
		return nil, err
	} else {
		if neededVersion <= schema.Version {
			health.Healthy = true
		} else {
			health.Reason = fmt.Sprintf("database version is not at least %d", neededVersion)
		}
		return health, nil
	}
}

