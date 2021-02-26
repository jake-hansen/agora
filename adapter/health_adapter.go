package adapter

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/domain"
)

func HealthDomainToDTO(health *domain.Health) *dto.Health {
	h := &dto.Health{
		Healthy: health.Healthy,
	}
	return h
}
