package meetingproviderhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/router/handlers"
)

func Provide(authMiddleware *authmiddleware.AuthMiddleware, meetingProviderService domain.MeetingProviderService) *MeetingProviderHandler {
	return &MeetingProviderHandler{
		AuthMiddleware: authMiddleware,
		MeetingProviderService: &meetingProviderService,
	}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*MeetingProviderHandler)))
)