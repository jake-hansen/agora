package simpleinviteservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(inviteRepo domain.InviteRepository, meetingService domain.MeetingPlatformService, oauthService domain.OAuthInfoService, userService domain.UserService) *SimpleInviteService {
	return &SimpleInviteService{
		inviteRepo:     inviteRepo,
		meetingService: meetingService,
		oauthService:   oauthService,
		userService:    userService,
	}
}

var (
	// ProviderProductionSet provides a new SimpleInviteService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.InviteService), new(*SimpleInviteService)))
)
