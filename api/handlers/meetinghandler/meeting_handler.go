package meetinghandler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"net/http"
)

type MeetingHandler struct {
	AuthMiddleware	*authmiddleware.AuthMiddleware
	PlatformService *domain.MeetingPlatformService
	OAuthService	*domain.OAuthInfoService
}

func (m *MeetingHandler) meetingPlatformValidator(c *gin.Context, platformName string) *domain.MeetingPlatform {
	platform, err := (*m.PlatformService).GetByPlatformName(platformName)
	if err != nil {
		apiError := api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("the platform %s does not exist", platformName))
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return nil
	}
	return platform
}

func (m *MeetingHandler) getUser(c *gin.Context) *domain.User {
	user, err := m.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return nil
	}
	return user
}

func (m *MeetingHandler) validateHelper(err error) error {
	var verr validator.ValidationErrors
	if err != nil && !errors.As(err, &verr) {
		err = api.NewAPIError(http.StatusBadRequest, err, "could not parse request body")
	}
	return err
}

func (m *MeetingHandler) Register(parentGroup *gin.RouterGroup) error {
	meetingHandlerGroup := parentGroup.Group("users")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.POST("/me/:platform/meetings", m.CreateMeeting)
		meetingHandlerGroup.GET("/me/:platform/meetings", m.GetMeetings)
	}

	return nil
}

func (m *MeetingHandler) CreateMeeting(c *gin.Context) {
	platformName := c.Param("platform")

	platform := m.meetingPlatformValidator(c, platformName)
	if platform == nil {
		return
	}

	user := m.getUser(c)
	if user == nil {
		return
	}

	var meeting dto.Meeting
	err := c.ShouldBind(&meeting)
	err = m.validateHelper(err)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	oauth, err := (*m.OAuthService).GetOAuthInfo(user.ID, platform)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	createdMeeting, err := platform.Actions.CreateMeeting(*oauth, adapter.MeetingDTOToDomain(&meeting))
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusCreated, adapter.MeetingDomainToDTO(createdMeeting))
}

func (m *MeetingHandler) GetMeetings(c *gin.Context)  {
	platformName := c.Param("platform")

	platform := m.meetingPlatformValidator(c, platformName)
	if platform == nil {
		return
	}

	user := m.getUser(c)
	if user == nil {
		return
	}

	oauth, err := (*m.OAuthService).GetOAuthInfo(user.ID, platform)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	meetings, err := platform.Actions.GetMeetings(*oauth)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.DomainMeetingPageToDTOMeetingPage(meetings))
}
