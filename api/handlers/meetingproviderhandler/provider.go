package meetingproviderhandler

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/router/handlers"
)

func Provide(authMiddleware *authmiddleware.AuthMiddleware) *MeetingProviderHandler {
	return &MeetingProviderHandler{
		AuthMiddleware: authMiddleware,
	}
}

var (
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(handlers.Handler), new(*MeetingProviderHandler)))
)