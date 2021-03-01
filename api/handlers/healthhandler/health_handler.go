package healthhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

// HealthHandler is the handler that manages operations on Health for the API.
type HealthHandler struct {
	healthService *domain.HealthService
}

// Register creates one endpoint to manage Health.
// / (GET) - Get current Health
func (h *HealthHandler) Register(parentGroup *gin.RouterGroup) error {
	healthGroup := parentGroup.Group("health")
	{
		healthGroup.GET("", h.GetHealth)
	}
	return nil
}

// GetHealth attempts to get the current Health and set it as a JSON response.
func (h *HealthHandler) GetHealth(c *gin.Context) {
	health, err := (*h.healthService).GetHealth()

	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	} else {
		if health.Healthy {
			c.JSON(http.StatusOK, adapter.HealthDomainToDTO(health))
		} else {
			c.JSON(http.StatusServiceUnavailable, adapter.HealthDomainToDTO(health))
		}
	}
}

