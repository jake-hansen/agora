package meetingproviderhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"net/http"
)

type MeetingProviderHandler struct {
	AuthMiddleware	*authmiddleware.AuthMiddleware
}

func (m *MeetingProviderHandler) Register(parentGroup *gin.RouterGroup) error {
	meetingHandlerGroup := parentGroup.Group("provider")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.POST(":meetingprovidername/auth", m.Auth)
	}

	return nil
}

func (m *MeetingProviderHandler) Auth(c *gin.Context) {
	meetingProviderName := c.Param("meetingprovidername")
	authorizationCode := c.Query("code")

	c.JSON(http.StatusOK, fmt.Sprintf("provider: %s, code: %s", meetingProviderName, authorizationCode))
}

