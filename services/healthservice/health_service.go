package healthservice

import (
	"errors"
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

var schemaVersion = 9

// HealthService is a service which processes information about the application's health.
type HealthService struct {
	schemaRepo *domain.SchemaMigrationRepo
}

// GetHealth retrieves the Health of the application.
func (h *HealthService) GetHealth() (*domain.Health, error) {
	return h.backwardsCompatibleVersionHealthCheck(schemaVersion)
}

// specificVersionNeededHealthCheck retrieves the Health of the application when a specific
// database schema version is needed.
func (h *HealthService) specificVersionNeededHealthCheck(neededVersion int) (*domain.Health, error) {
	health := &domain.Health{Healthy: false}

	_, err := (*h.schemaRepo).GetSchemaMigrationByVersion(neededVersion)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			health.Reason = fmt.Sprintf("database version is not at needed version %d", schemaVersion)
		} else {
			return nil, err
		}
	} else {
		health.Healthy = true
	}

	return health, nil
}

// backwardsCompatibleVersionHealthCheck retrieves the Health of the application when a
// backwards compatible database schema version can be used.
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

