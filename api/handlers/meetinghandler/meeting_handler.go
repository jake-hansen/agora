package meetinghandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/zoom"
)

type MeetingHandler struct {
	AuthMiddleware  *authmiddleware.AuthMiddleware
	PlatformService *domain.MeetingPlatformService
	OAuthService    *domain.OAuthInfoService
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
		meetingHandlerGroup.GET("/me/:platform/meetings/:id", m.GetMeeting)
	}

	return nil
}

func (m *MeetingHandler) platformErrorConverter(err error) error {
	var apiErr = err
	if errors.Is(err, zoom.ErrReqCreation) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while formulating the request. Please try again later.")
	} else if errors.Is(err, zoom.ErrReqExecution) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while executing the request. Please try again later.")
	} else if errors.Is(err, zoom.ErrResDecoding) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while decoding the performed request. Please try again later.")
	}
	return apiErr
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
		err = m.platformErrorConverter(err)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusCreated, adapter.MeetingDomainToDTO(createdMeeting))
}

func (m *MeetingHandler) GetMeetings(c *gin.Context) {
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
		err = m.platformErrorConverter(err)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.DomainMeetingPageToDTOMeetingPage(meetings))
}

func (m *MeetingHandler) GetMeeting(c *gin.Context) {
	platformName := c.Param("platform")
	meetingID := c.Param("id")

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

	meeting, err := platform.Actions.GetMeeting(*oauth, meetingID)
	if err != nil {
		err = m.platformErrorConverter(err)
		if errors.Is(err, zoom.ErrNotFound) {
			err = api.NewAPIError(http.StatusNotFound, err, "the requested meeting was not found")
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.MeetingDomainToDTO(meeting))
}
