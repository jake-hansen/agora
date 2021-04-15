package invitehandler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/database/repositories"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/platforms/common"
	"github.com/jake-hansen/agora/services/simpleinviteservice"
	"net/http"
	"strconv"
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
		inviteGroup.DELETE("/:id", i.DeleteInvite)
	}
	
	userGroup := parentGroup.Group("/users")
	userGroup.Use(i.AuthMiddleware.HandleAuth())
	{
		userGroup.GET("/:id/invites", i.GetInvites)
	}

	return nil
}

func (i *InviteHandler) SendInvite(c *gin.Context) {
	var invite dto.InviteRequest
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

	if user.ID != invite.InviterID {
		err = errors.New("inviter ID not same as authenticated user ID")
		apiErr := api.NewAPIError(http.StatusBadRequest, err, "cannot send invite for different user")
		_ = c.Error(apiErr).SetType(gin.ErrorTypePublic)
		return
	}

	inviteID, err := (*i.InviteService).SendInvite(adapter.InviteRequestDTOToDomain(&invite))
	if err != nil {
		var notFoundError repositories.NotFoundError
		var meetingNotFoundError common.NotFoundError
		var duplicateEntry repositories.DuplicateEntryError

		if errors.As(err, &simpleinviteservice.InviterSameAsInviteeErr{}) {
			err = api.NewAPIError(http.StatusBadRequest, err, "inviter cannot be invitee")
		} else if errors.As(err, &notFoundError) {
			err = api.NewAPIError(http.StatusBadRequest, err, fmt.Sprintf("'%s' was not found", notFoundError.Value))
		} else if errors.As(err, &meetingNotFoundError) {
			err = api.NewAPIError(http.StatusBadRequest, err, fmt.Sprintf("meeting with id %s not found", invite.MeetingID))
		} else if errors.As(err, &duplicateEntry) {
			err = api.NewAPIError(http.StatusBadRequest, err, fmt.Sprintf("cannot create %s that already exists", duplicateEntry.EntityType))
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resource := dto.Resource{ID: int(inviteID)}

	c.JSON(http.StatusCreated, resource)
}

func (i *InviteHandler) GetInvites(c *gin.Context)  {
	inviteType := c.Query("type")
	if inviteType != "sent" {
		if inviteType != "received" {
			err := errors.New("unknown type query")
			apiErr := api.NewAPIError(http.StatusBadRequest, err, "'type' query must be either 'sent' or 'received'")
			_ = c.Error(apiErr).SetType(gin.ErrorTypePublic)
			return
		}
	}

	user, err := i.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	var invites []*domain.Invite

	if inviteType == "sent" {
		invites, err = (*i.InviteService).GetAllSentInvites(user.ID)
		if err != nil {
			return
		}
	}

	if inviteType == "received" {
		invites, err = (*i.InviteService).GetAllReceivedInvites(user.ID)
		if err != nil {
			return
		}
	}

	var dtoInvites []*dto.Invite
	for _, invite := range invites {
		dtoInvites = append(dtoInvites, adapter.InviteDomainToDTO(invite))
	}

	c.JSON(http.StatusOK, dtoInvites)
}

func (i *InviteHandler) DeleteInvite(c *gin.Context)  {
	inviteIDStr := c.Param("id")
	inviteID, err := strconv.Atoi(inviteIDStr)
	if err != nil {
		apiErr := api.NewAPIError(http.StatusBadRequest, err, "invite id must be a parsable integer")
		_ = c.Error(apiErr).SetType(gin.ErrorTypePublic)
		return
	}

	user, err := i.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	invite, err := (*i.InviteService).GetInvite(uint(inviteID))
	if err != nil {
		var notFoundErr repositories.NotFoundError
		if  errors.As(err, &notFoundErr){
			err = api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("invite with id %s not found", inviteIDStr))
		}
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	if invite.InviterID != user.ID {
		err = errors.New("cannot delete invite created by another user")
		apiErr := api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("invite with id %s not found", inviteIDStr))
		_ = c.Error(apiErr).SetType(gin.ErrorTypePublic)
		return
	}

	err = (*i.InviteService).DeleteInvite(uint(inviteID))
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
