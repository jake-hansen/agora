package meetingplatformhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

func Provide(authMiddleware *authmiddleware.AuthMiddleware, meetingProviderService domain.MeetingPlatformService, oauthService domain.OAuthInfoService) *MeetingPlatformHandler {
	return &MeetingPlatformHandler{
		AuthMiddleware:  authMiddleware,
		PlatformService: &meetingProviderService,
		OAuthService: &oauthService,
	}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*MeetingPlatformHandler)))
)