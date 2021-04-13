// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package loader

import (
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/platforms"
	"github.com/jake-hansen/agora/platforms/webex"
	"github.com/jake-hansen/agora/platforms/zoom"
	"github.com/spf13/viper"
)

// Injectors from injector.go:

// Build returns a Loader configured with the provided Manager and Viper.
func Build(db *database.Manager, v *viper.Viper) (*Loader, error) {
	meetingPlatformRepo := meetingplatformrepo.Provide(db)
	zoomActions := zoom.Provide()
	webexActions := webex.Provide()
	configuredPlatforms := platforms.Provide(zoomActions, webexActions, v)
	meetingPlatformLoader := ProvideMeetingPlatformLoader(meetingPlatformRepo, configuredPlatforms)
	loader := ProvideLoader(meetingPlatformLoader)
	return loader, nil
}
