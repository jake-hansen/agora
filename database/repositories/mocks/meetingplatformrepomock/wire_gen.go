// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package meetingplatformrepomock

// Injectors from injector.go:

func Build() *MeetingPlatformRepository {
	meetingProviderRepository := Provide()
	return meetingProviderRepository
}
