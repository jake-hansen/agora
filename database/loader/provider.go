package loader

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/providers"
)

func ProvideLoader(loader *MeetingPlatformLoader) *Loader {
	return NewLoader(loader)
}

func ProvideMeetingPlatformLoader(repo domain.MeetingPlatformRepository, configuredPlatforms []*domain.MeetingPlatform) *MeetingPlatformLoader {
	return NewMeetingPlatformLoader(repo, configuredPlatforms)
}

var (
	ProviderProductionSet = wire.NewSet(ProvideLoader, ProvideMeetingPlatformLoader, providers.ProductionSet)
)
