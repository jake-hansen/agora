package healthhandler

import "github.com/jake-hansen/agora/domain"

func Provide(healthService domain.HealthService) *HealthHandler {
	return &HealthHandler{healthService: &healthService}
}
