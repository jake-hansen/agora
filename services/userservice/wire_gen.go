// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package userservice

import (
	"github.com/jake-hansen/agora/config"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/userrepo"
	"github.com/jake-hansen/agora/domain"
)

// Injectors from injector.go:

func Build() (*UserService, func(), error) {
	viper := config.Provide()
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
	userService := Provide(userRepository)
	return userService, func() {
		cleanup()
	}, nil
}

func BuildTest(repo domain.UserRepository) *UserService {
	userService := Provide(repo)
	return userService
}