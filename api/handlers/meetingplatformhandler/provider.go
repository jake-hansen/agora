package meetingplatformhandler

import (
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a MeetingPlatformHandler configured with the provided AuthMiddleware, MeetingPlatformService, and OAuthService.
func Provide(authMiddleware *authmiddleware.AuthMiddleware, meetingProviderService domain.MeetingPlatformService, oauthService domain.OAuthInfoService) *MeetingPlatformHandler {
	return &MeetingPlatformHandler{
		AuthMiddleware:  authMiddleware,
		PlatformService: &meetingProviderService,
		OAuthService:    &oauthService,
	}
}
