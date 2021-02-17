package meetingproviderhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

type MeetingProviderHandler struct {
	AuthMiddleware	*authmiddleware.AuthMiddleware
	Providers		[]domain.MeetingProvider
}

func (m *MeetingProviderHandler) Register(parentGroup *gin.RouterGroup) error {
	meetingHandlerGroup := parentGroup.Group("provider")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.GET("", m.GetAllProviders)
		meetingHandlerGroup.POST(":meetingprovidername/auth", m.Auth)
	}

	return nil
}

func (m *MeetingProviderHandler) Auth(c *gin.Context) {
	meetingProviderName := c.Param("meetingprovidername")
	authorizationCode := c.Query("code")

	c.JSON(http.StatusOK, fmt.Sprintf("provider: %s, code: %s", meetingProviderName, authorizationCode))
}

func (m *MeetingProviderHandler) GetAllProviders(c *gin.Context) {
	var providers []dto.MeetingProvider

	for _, provider := range m.Providers {
		providers = append(providers, *adapter.MeetingProviderDomainToDTO(provider))
	}
	c.JSON(http.StatusOK, providers)
}

