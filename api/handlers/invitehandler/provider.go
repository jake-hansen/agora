package invitehandler

import (
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
)

func Provide(inviteService domain.InviteService, authMiddleware *authmiddleware.AuthMiddleware, userService domain.UserService, platformService domain.MeetingPlatformService, oauthService domain.OAuthInfoService) *InviteHandler {
	return &InviteHandler{
		InviteService:   &inviteService,
		AuthMiddleware:  authMiddleware,
		UserService:     &userService,
		PlatformService: &platformService,
		OAuthService:    &oauthService,
	}
}