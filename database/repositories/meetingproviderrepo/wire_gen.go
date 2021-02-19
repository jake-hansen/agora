// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package meetingproviderrepo

import (
	"github.com/jake-hansen/agora/database"
)

// Injectors from injector.go:

func Build(manager *database.Manager) (*MeetingProviderRepo, func(), error) {
	meetingProviderRepo := Provide(manager)
	return meetingProviderRepo, func() {
	}, nil
}
