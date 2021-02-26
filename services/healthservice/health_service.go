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
	health := &domain.Health{Healthy: false}

	_, err := (*h.schemaRepo).GetSchemaMigrationByVersion(neededSchemaVersion)
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

