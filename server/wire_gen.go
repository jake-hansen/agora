// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package server

import (
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/handlers/authhandler"
	"github.com/jake-hansen/agora/api/handlers/healthhandler"
	"github.com/jake-hansen/agora/api/handlers/invitehandler"
	"github.com/jake-hansen/agora/api/handlers/meetinghandler"
	"github.com/jake-hansen/agora/api/handlers/meetingplatformhandler"
	"github.com/jake-hansen/agora/api/handlers/userhandler"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/api/middleware/authmiddleware"
	"github.com/jake-hansen/agora/api/middleware/corsmiddleware"
	"github.com/jake-hansen/agora/api/validator"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/inviterepo"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/database/repositories/oauthinforepo"
	"github.com/jake-hansen/agora/database/repositories/refreshtokenrepo"
	"github.com/jake-hansen/agora/database/repositories/schemamigrationrepo"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/log"
	"github.com/jake-hansen/agora/platforms"
	"github.com/jake-hansen/agora/platforms/webex"
	"github.com/jake-hansen/agora/platforms/zoom"
	"github.com/jake-hansen/agora/router"
	handlers2 "github.com/jake-hansen/agora/router/handlers"
	"github.com/jake-hansen/agora/services/cookieservice"
	"github.com/jake-hansen/agora/services/healthservice"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/meetingplatformservice"
	"github.com/jake-hansen/agora/services/oauthinfoservice"
	"github.com/jake-hansen/agora/services/refreshtokenservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/jake-hansen/agora/services/simpleinviteservice"
	"github.com/jake-hansen/agora/services/userservice"
	"github.com/spf13/viper"
)

// Injectors from injector.go:

func Build(db *database.Manager, v *viper.Viper, log2 *log.Log) (*Server, error) {
	config, err := Cfg(v)
	if err != nil {
		return nil, err
	}
	corsConfig, err := corsmiddleware.Cfg(v)
	if err != nil {
		return nil, err
	}
	corsMiddleware := corsmiddleware.Provide(corsConfig)
	v2 := middleware.ProvideAllProductionMiddleware(log2, corsMiddleware)
	jwtserviceConfig, err := jwtservice.Cfg(v)
	if err != nil {
		return nil, err
	}
	jwtServiceImpl := jwtservice.Provide(jwtserviceConfig)
	userRepository := userrepo.Provide(db)
	userService := userservice.Provide(userRepository)
	refreshTokenRepo := refreshtokenrepo.Provide(db)
	refreshTokenService := refreshtokenservice.Provide(refreshTokenRepo)
	simpleAuthService := simpleauthservice.Provide(jwtServiceImpl, userService, refreshTokenService)
	cookieserviceConfig := cookieservice.Cfg(v)
	cookieService := cookieservice.Provide(cookieserviceConfig)
	v3 := authmiddleware.ProvideAuthorizationHeaderParser()
	authMiddleware := authmiddleware.Provide(simpleAuthService, v3)
	authHandler := authhandler.Provide(simpleAuthService, cookieService, authMiddleware)
	userHandler := userhandler.Provide(userService, authMiddleware)
	meetingPlatformRepo := meetingplatformrepo.Provide(db)
	zoomActions := zoom.Provide()
	webexActions := webex.Provide()
	configuredPlatforms := platforms.Provide(zoomActions, webexActions, v)
	meetingPlatformService := meetingplatformservice.Provide(meetingPlatformRepo, configuredPlatforms)
	oAuthInfoRepo := oauthinforepo.Provide(db)
	oAuthInfoService := oauthinfoservice.Provide(meetingPlatformService, oAuthInfoRepo)
	meetingPlatformHandler := meetingplatformhandler.Provide(authMiddleware, meetingPlatformService, oAuthInfoService)
	schemaMigrationRepo := schemamigrationrepo.Provide(db)
	healthService := healthservice.Provide(schemaMigrationRepo)
	healthHandler := healthhandler.Provide(healthService)
	inviteRepo := inviterepo.Provide(db)
	simpleInviteService := simpleinviteservice.Provide(inviteRepo, meetingPlatformService, oAuthInfoService, userService)
	meetingHandler := meetinghandler.Provide(authMiddleware, meetingPlatformService, oAuthInfoService, simpleInviteService)
	inviteHandler := invitehandler.Provide(simpleInviteService, authMiddleware, userService, meetingPlatformService, oAuthInfoService)
	v4 := handlers.ProvideAllProductionHandlers(authHandler, userHandler, meetingPlatformHandler, healthHandler, meetingHandler, inviteHandler)
	handlerManager := handlers2.ProvideHandlerManager(v4)
	v5 := validator.ProvideCustomValidationFuncs()
	validatorConfig := validator.Cfg(v5)
	validatorValidator, err := validator.Provide(validatorConfig)
	if err != nil {
		return nil, err
	}
	routerConfig, err := router.Cfg(v, v2, handlerManager, validatorValidator)
	if err != nil {
		return nil, err
	}
	routerRouter := router.Provide(routerConfig)
	server := Provide(config, routerRouter)
	return server, nil
}
