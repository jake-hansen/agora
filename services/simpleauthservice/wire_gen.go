// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package simpleauthservice

import (
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/servicemocks/userservicemock"
	"github.com/jake-hansen/agora/services/userservice"
)

// Injectors from injector.go:

func Build() (*SimpleAuthService, func(), error) {
	viper := config.Provide()
	jwtserviceConfig, err := jwtservice.Cfg(viper)
	if err != nil {
		return nil, nil, err
	}
	service := jwtservice.Provide(jwtserviceConfig)
	databaseConfig, err := database.Cfg(viper)
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := database.ProvideGORM(databaseConfig)
	if err != nil {
		return nil, nil, err
	}
	manager, err := database.Provide(databaseConfig, db)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepository := userrepo.Provide(manager)
	userService := userservice.Provide(userRepository)
	simpleAuthService := Provide(service, userService)
	return simpleAuthService, func() {
		cleanup()
	}, nil
}

func BuildTest(jwtCfg jwtservice.Config) (*SimpleAuthService, error) {
	jwtserviceConfig, err := jwtservice.CfgTest(jwtCfg)
	if err != nil {
		return nil, err
	}
	service := jwtservice.Provide(jwtserviceConfig)
	userService := userservicemock.Provide()
	simpleAuthService := Provide(service, userService)
	return simpleAuthService, nil
}
