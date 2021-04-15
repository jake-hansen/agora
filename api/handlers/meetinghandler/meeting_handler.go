package meetinghandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jake-hansen/agora/platforms/common"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

// MeetingHandler is the handler that manages operations on Meetings for the API.
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

// Register creates 3 endpoints to manage meetings
// /me/platforms/:platform/meetings (GET) - retrieves all meetings for the authenticated user on the specified platform
// /me/platforms/:platform/meetings (POST) - creates a new meeting for the authenticated user on the specified platform
// /me/platforms/:platform/meetings/:id (GET) - retrieves the specified meeting for the authenticated user on the specified platform
func (m *MeetingHandler) Register(parentGroup *gin.RouterGroup) error {
	meetingHandlerGroup := parentGroup.Group("users")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.POST("/:id/platforms/:platform/meetings", m.CreateMeeting)
		meetingHandlerGroup.GET("/:id/platforms/:platform/meetings", m.GetMeetings)
		meetingHandlerGroup.GET("/:id/platforms/:platform/meetings/:id", m.GetMeeting)
		meetingHandlerGroup.DELETE("/:id/platforms/:platform/meetings/:id", m.DeleteMeeting)
	}

	return nil
}

func (m *MeetingHandler) platformErrorConverter(err error) error {
	var apiErr = err
	if errors.Is(err, common.ErrReqCreation) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while formulating the request. Please try again later.")
	} else if errors.Is(err, common.ErrReqExecution) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while executing the request. Please try again later.")
	} else if errors.Is(err, common.ErrResDecoding) {
		apiErr = api.NewAPIError(http.StatusInternalServerError, err,
			"An error occurred while decoding the performed request. Please try again later.")
	}
	return apiErr
}

// CreateMeeting creates a meeting
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

	meetingTypeString := c.Query("type")
	meetingType, _ := strconv.Atoi(meetingTypeString)
	if !(meetingType == 1 || meetingType == 2) {
		apiErr := api.NewAPIError(http.StatusBadRequest, errors.New("bad request"), "meeting type should be 1 (instant) or 2 (scheduled)")
		_ = c.Error(apiErr).SetType(gin.ErrorTypePublic)
		return
	}

	var scheduledMeeting dto.Meeting
	var instantMeeting dto.InstantMeeting
	var domainMeeting *domain.Meeting
	var err error
	if meetingType == 1 {
		err = c.ShouldBind(&instantMeeting)
		domainMeeting = adapter.InstantMeetingDTOToDomain(&instantMeeting)
	} else if meetingType == 2 {
		err = c.ShouldBind(&scheduledMeeting)
		domainMeeting = adapter.ScheduledMeetingDTOToDomain(&scheduledMeeting)
	}
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

	createdMeeting, err := platform.Actions.CreateMeeting(*oauth, domainMeeting)
	if err != nil {
		err = m.platformErrorConverter(err)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusCreated, adapter.MeetingDomainToDTO(createdMeeting))
}

// GetMeeting gets all meetings
func (m *MeetingHandler) GetMeetings(c *gin.Context) {
	platformName := c.Param("platform")
	pageSize := c.Query("page_size")
	nextPage := c.Query("next_page")

	var size int
	var err error
	if pageSize != "" {
		size, err = strconv.Atoi(pageSize)
		if err != nil {
			err = api.NewAPIError(http.StatusBadRequest, err, "could not parse page_size")
			_ = c.Error(err).SetType(gin.ErrorTypePublic)
			return
		}
	}

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

	pageReq := dto.NewPageReq(size, nextPage)

	meetings, err := platform.Actions.GetMeetings(*oauth, *adapter.PageRequestDTOToDomain(pageReq))
	if err != nil {
		err = m.platformErrorConverter(err)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.MeetingPageDomainToDTO(meetings))
}

// GetMeeting gets a single meeting
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
		if errors.Is(err, common.ErrNotFound) {
			err = api.NewAPIError(http.StatusNotFound, err, "the requested meeting was not found")
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, adapter.MeetingDomainToDTO(meeting))
}

func (m *MeetingHandler) DeleteMeeting(c *gin.Context) {
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

	err = platform.Actions.DeleteMeeting(*oauth, meetingID)
	if err != nil {
		var notFoundErr common.NotFoundError

		if errors.Is(err, notFoundErr) {
			err = api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("meeting with id %s not found", meetingID))
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
