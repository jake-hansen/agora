package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/router/handlers"
)

func ProvideAllProductionHandlers() ([]handlers.Handler, func(), error) {
	var handlers []handlers.Handler

	auth, authCleanup, err := authhandler.Build()
	if err != nil {
		return nil, nil, err
	}
	user, userCleanup, err := userhandler.Build()
	if err != nil {
		return nil, nil, err
	}

	handlers = append(handlers, auth)
	handlers = append(handlers, user)

	cleanup := func() {
		authCleanup()
		userCleanup()
	}

	return handlers, cleanup, nil
}

var (
	ProviderProductionSet = wire.NewSet(ProvideAllProductionHandlers)
)
