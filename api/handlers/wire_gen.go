// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package handlers

import (
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/log"
	"github.com/jake-hansen/agora/router/handlers"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/userservice"
)

// Injectors from injector.go:

func Build() (*[]handlers.Handler, func(), error) {
	viper := config.Provide()
	jwtserviceConfig, err := jwtservice.Cfg(viper)
	if err != nil {
		return nil, nil, err
	}
	jwtServiceImpl := jwtservice.Provide(jwtserviceConfig)
	zapConfig := log.Cfg(viper)
	logLog, cleanup, err := log.Provide(zapConfig)
	if err != nil {
		return nil, nil, err
	}
	databaseConfig, err := database.Cfg(viper, logLog)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup2, err := database.ProvideGORM(databaseConfig)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	manager, err := database.Provide(databaseConfig, db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	userRepository := userrepo.Provide(manager)
	userService := userservice.Provide(userRepository)
	simpleAuthService := simpleauthservice.Provide(jwtServiceImpl, userService)
	authHandler := authhandler.Provide(simpleAuthService)
	userHandler := userhandler.Provide(userService)
	v := ProvideAllProductionHandlers(authHandler, userHandler)
	return v, func() {
		cleanup2()
		cleanup()
	}, nil
}
