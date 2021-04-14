package invitehandler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/common"
	"net/http"
)

type InviteHandler struct {
	InviteService *domain.InviteService
	AuthMiddleware *authmiddleware.AuthMiddleware
	UserService *domain.UserService
	PlatformService *domain.MeetingPlatformService
	OAuthService *domain.OAuthInfoService
}

func (i *InviteHandler) platformErrorConverter(err error) error {
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

func (i *InviteHandler) meetingPlatformValidator(c *gin.Context, platformName string) *domain.MeetingPlatform {
	platform, err := (*i.PlatformService).GetByPlatformName(platformName)
	if err != nil {
		apiError := api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("the platform %s does not exist", platformName))
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return nil
	}
	return platform
}

func (i *InviteHandler) Register(parentGroup *gin.RouterGroup) error {
	inviteGroup := parentGroup.Group("/invites")
	inviteGroup.Use(i.AuthMiddleware.HandleAuth())
	{
		inviteGroup.POST("", i.SendInvite)
	}

	return nil
}

func (i *InviteHandler) SendInvite(c *gin.Context) {
	var invite dto.Invite
	err := c.ShouldBind(&invite)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := i.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Validate platform
	platform := i.meetingPlatformValidator(c, invite.MeetingPlatform)
	if platform == nil {
		return
	}

	// Validate meeting exists
	oauth, err := (*i.OAuthService).GetOAuthInfo(user.ID, platform)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	meeting, err := platform.Actions.GetMeeting(*oauth, invite.MeetingID)
	if err != nil {
		err = i.platformErrorConverter(err)
		if errors.Is(err, common.ErrNotFound) {
			err = api.NewAPIError(http.StatusNotFound, err, "the requested meeting was not found")
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	// Validate invitee exists
	invitee, err := (*i.UserService).GetByUsername(invite.Invitee)
	if err != nil {
		return
	}

	// Create invite
	domainInvite := &domain.Invite{
		MeetingID:      invite.MeetingID,
		MeetingEndTime: meeting.StartTime.Add(meeting.Duration),
		InviterID:      invite.InviterID,
		InviteeID:      invitee.ID,
	}

	inviteID, err := (*i.InviteService).SendInvite(domainInvite)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resource := dto.Resource{ID: int(inviteID)}

	c.JSON(http.StatusCreated, resource)
}
