package healthhandler

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

type HealthHandler struct {
	healthService *domain.HealthService
}

func (h *HealthHandler) Register(parentGroup *gin.RouterGroup) error {
	healthGroup := parentGroup.Group("health")
	{
		healthGroup.GET("", h.GetHealth)
	}
	return nil
}

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

