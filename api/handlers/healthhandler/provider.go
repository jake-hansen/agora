package healthhandler

import "github.com/jake-hansen/agora/domain"

// Provide returns a HealthHandler configured with the given HealthService.
func Provide(healthService domain.HealthService) *HealthHandler {
	return &HealthHandler{healthService: &healthService}
}
