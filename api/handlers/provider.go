package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/router/handlers"
)

func ProvideAllProductionHandlers() []handlers.Handler {
	var handlers []handlers.Handler

	auth, _, _ := authhandler.Build()

	handlers = append(handlers, auth)

	return handlers
}

var (
	ProviderProductionSet = wire.NewSet(ProvideAllProductionHandlers)
)
