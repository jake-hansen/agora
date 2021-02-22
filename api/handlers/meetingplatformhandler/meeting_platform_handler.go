package meetingplatformhandler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/adapter"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/oauth2"
	"net/http"
)

type MeetingPlatformHandler struct {
	AuthMiddleware  *authmiddleware.AuthMiddleware
	PlatformService *domain.MeetingPlatformService
	OAuthService	*domain.OAuthInfoService
}

func (m *MeetingPlatformHandler) Register(parentGroup *gin.RouterGroup) error {
	meetingHandlerGroup := parentGroup.Group("platform")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.GET("", m.GetAllPlatforms)
		meetingHandlerGroup.POST(":platform/auth", m.Auth)
		meetingHandlerGroup.GET(":platform/auth", m.GetAuth)
	}

	return nil
}

func (m *MeetingPlatformHandler) meetingPlatformValidator(c *gin.Context, platformName string) *domain.MeetingPlatform {
	platform, err := (*m.PlatformService).GetByPlatformName(platformName)
	if err != nil {
		apiError := api.NewAPIError(http.StatusNotFound, err, fmt.Sprintf("the platform %s does not exist", platformName))
		_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
		return nil
	}
	return platform
}

func (m *MeetingPlatformHandler) getUser(c *gin.Context) *domain.User {
	user, err := m.AuthMiddleware.GetUser(c)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return nil
	}
	return user
}

func (m *MeetingPlatformHandler) Auth(c *gin.Context) {
	platformName := c.Param("platform")
	authorizationCode := c.Query("code")

	platform := m.meetingPlatformValidator(c, platformName)
	if platform == nil {
		return
	} else {
		user := m.getUser(c)
		if user == nil {
			return
		}
		err := (*m.OAuthService).CreateOAuthInfo(c, authorizationCode, user.ID, platform)
		if err != nil {
			var apiError *api.APIError
			var oauthError *oauth2.RetrieveError
			var tokenExistsError *domain.TokenExistsError
			if errors.As(err, &oauthError) {
				apiError = api.NewAPIError(http.StatusBadRequest, oauthError, "could not validate authorization code")
			} else if errors.As(err, &tokenExistsError) {
				apiError = api.NewAPIError(http.StatusBadRequest, tokenExistsError, "authentication tokens already exist for this platform")
			} else {
				apiError = api.NewAPIError(http.StatusInternalServerError, err, "an error occurred while saving the authentication tokens")
			}
			_ = c.Error(apiError).SetType(gin.ErrorTypePublic)
			return
		}
	}

	c.Status(http.StatusOK)
}

func (m *MeetingPlatformHandler) GetAuth(c *gin.Context) {
	platformName := c.Param("platform")

	user := m.getUser(c)
	if user == nil {
		return
	}

	platform := m.meetingPlatformValidator(c, platformName)
	if platform == nil {
		return
	} else {
		_, err := (*m.OAuthService).GetOAuthInfo(user.ID, platform)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func (m *MeetingPlatformHandler) GetAllPlatforms(c *gin.Context) {
	var platforms []dto.MeetingProvider
	var retrievedPlatforms []*domain.MeetingPlatform
	retrievedPlatforms, err := (*m.PlatformService).GetAll()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	for _, provider := range retrievedPlatforms {
		platforms = append(platforms, *adapter.MeetingProviderDomainToDTO(*provider))
	}
	c.JSON(http.StatusOK, platforms)
}

