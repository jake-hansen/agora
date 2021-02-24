package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/handlers/meetingplatformhandler"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/database/repositories/oauthinforepo"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/router/handlers"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/meetingplatforms"
	"github.com/jake-hansen/agora/services/meetingplatformservice"
	"github.com/jake-hansen/agora/services/oauthinfoservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/userservice"
)

// ProvideAllProductionHandlers provides all the handlers that will be used in production.
func ProvideAllProductionHandlers(auth *authhandler.AuthHandler, user *userhandler.UserHandler,
		meetingProvider *meetingplatformhandler.MeetingPlatformHandler) *[]handlers.Handler {

	var handlers []handlers.Handler

	handlers = append(handlers, auth)
	handlers = append(handlers, user)
	handlers = append(handlers, meetingProvider)

	return &handlers
}

var (
	services = wire.NewSet(simpleauthservice.ProviderProductionSet,
		meetingplatformservice.ProviderSet,
		jwtservice.ProviderProductionSet,
		oauthinfoservice.ProviderProductionSet,
		userservice.ProviderProductionSet)

	repos = wire.NewSet(meetingplatformrepo.ProviderSet,
		meetingplatforms.ProviderSet,
		userrepo.ProviderProductionSet,
		oauthinforepo.ProviderSet)

	middleware = wire.NewSet(authmiddleware.Provide,
		authmiddleware.ProvideAuthorizationHeaderParser)

	handlersSet = wire.NewSet(authhandler.Provide,
		userhandler.Provide,
		meetingplatformhandler.Provide)

	// ProviderProductionSet provides all handlers for production.
	ProviderProductionSet = wire.NewSet(ProvideAllProductionHandlers, repos, services, middleware, handlersSet)
)
