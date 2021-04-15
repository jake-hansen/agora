package invitehandler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
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
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	resource := dto.Resource{ID: int(inviteID)}

	c.JSON(http.StatusCreated, resource)
}

func (i *InviteHandler) GetInvites(c *gin.Context)  {
	user, err := i.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	invites, err := (*i.InviteService).GetAllReceivedInvites(user.ID)
	if err != nil {
		return 
	}

	var dtoInvites []*dto.Invite
	for _, invite := range invites {
		dtoInvites = append(dtoInvites, adapter.InviteDomainToDTO(invite))
	}

	c.JSON(http.StatusOK, dtoInvites)
}
