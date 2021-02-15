package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/providers"
	"github.com/jake-hansen/agora/router/handlers"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/userservice"
)

// ProvideAllProductionHandlers provides all the handlers that will be used in production.
func ProvideAllProductionHandlers(auth *authhandler.AuthHandler, user *userhandler.UserHandler) *[]handlers.Handler {
	var handlers []handlers.Handler

	handlers = append(handlers, auth)
	handlers = append(handlers, user)

	return &handlers
}

var (
	authHandlerProductionSet = wire.NewSet(authhandler.Provide,
		simpleauthservice.ProviderProductionSet,
		userservice.ProviderProductionSet,
		userrepo.ProviderProductionSet,
		jwtservice.ProviderProductionSet)

	userHandlerProductionSet = wire.NewSet(userhandler.Provide)

	// ProviderProductionSet provides all handlers for production.
	ProviderProductionSet = wire.NewSet(ProvideAllProductionHandlers,
		authHandlerProductionSet,
		userHandlerProductionSet,
		database.ProviderProductionSet,
		providers.ProductionSet)
)
