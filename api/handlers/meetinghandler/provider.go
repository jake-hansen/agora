package meetinghandler

import (
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a MeetingHandler configured with the provided AuthMiddleware, MeetingPlatformService, OAuthInfoService, and InviteService.
func Provide(authMiddleware *authmiddleware.AuthMiddleware, meetingProviderService domain.MeetingPlatformService, oauthService domain.OAuthInfoService, inviteService domain.InviteService) *MeetingHandler {
	return &MeetingHandler{
		AuthMiddleware:  authMiddleware,
		PlatformService: &meetingProviderService,
		OAuthService:    &oauthService,
		InviteService:   &inviteService,
	}
}
