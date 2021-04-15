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

// MeetingPlatformHandler is the handler that manages operations on MeetingPlatforms for the API.
type MeetingPlatformHandler struct {
	AuthMiddleware  *authmiddleware.AuthMiddleware
	PlatformService *domain.MeetingPlatformService
	OAuthService	*domain.OAuthInfoService
}

// Register creates three endpoints to manage Health.
// / 			  (GET)  - Gets all available MeetingPlatforms
// :platform/auth (POST) - Attempts to authenticate to the specified MeetingPlatform
// :platform/auth (GET)	 - Attempts to get the Auth for the specified MeetingPlatform
func (m *MeetingPlatformHandler) Register(parentGroup *gin.RouterGroup) error {
	userHandlerGroup := parentGroup.Group("users")
	userHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		userHandlerGroup.GET("/:id/platforms", m.GetAllAuth)
	}
	meetingHandlerGroup := parentGroup.Group("platforms")
	meetingHandlerGroup.Use(m.AuthMiddleware.HandleAuth())
	{
		meetingHandlerGroup.GET("", m.GetAllPlatforms)
		meetingHandlerGroup.POST("/:platform/auth", m.Auth)
		meetingHandlerGroup.GET("/:platform/auth", m.GetAuth)
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

// Auth attempts to authenticate a user against a MeetingPlatform using the
// provide authorization code. A 200 OK status is returned if OAuth tokens
// were successfully retrieved from the MeetingPlatform for the user. A
// 400 BAD REQUEST response is returned if the user already has OAuth tokens
// stored for the MeetingPlatform or if the provided authorization code was not
// validated by the platform. A 500 INTERNAL SERVER ERROR is returned if an error
// occurs.
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

// GetAuth attempts to get find an OAuthInfo object for the user that sent
// the request for the provided platform. If an OAuthInfo object exists,
// a 200 OK status is returned, otherwise a 404 NOT FOUND status is returned.
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

func (m *MeetingPlatformHandler) GetAllAuth(c *gin.Context) {
	user := m.getUser(c)
	if user == nil {
		return
	}
	platforms, err := (*m.OAuthService).GetAllAuthenticatedPlatforms(user.ID)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	var returnPlatforms []*dto.MeetingPlatform

	for _, platform := range platforms {
		returnPlatforms = append(returnPlatforms, adapter.MeetingPlatformDomainToDTO(*platform))
	}

	c.JSON(http.StatusOK, returnPlatforms)
}

// GetAllPlatforms retrieves all configured MeetingPlatforms and
// returns them as a JSON response.
func (m *MeetingPlatformHandler) GetAllPlatforms(c *gin.Context) {
	var platforms []dto.MeetingPlatform
	var retrievedPlatforms []*domain.MeetingPlatform
	retrievedPlatforms, err := (*m.PlatformService).GetAll()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	for _, provider := range retrievedPlatforms {
		platforms = append(platforms, *adapter.MeetingPlatformDomainToDTO(*provider))
	}
	c.JSON(http.StatusOK, platforms)
}
