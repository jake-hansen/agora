package loader

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// ProvideLoader provides a Loader configured with the provided MeetingPlatformLoader.
func ProvideLoader(loader *MeetingPlatformLoader) *Loader {
	return NewLoader(loader)
}

// ProvideMeetingPlatformLoader provides a MeetingPlatformLoader configured with the provided
// MeetingPlatformReposiotry and ConfiguredPlatforms.
func ProvideMeetingPlatformLoader(repo domain.MeetingPlatformRepository, configuredPlatforms domain.ConfiguredPlatforms) *MeetingPlatformLoader {
	return NewMeetingPlatformLoader(repo, configuredPlatforms)
}

var (
	// ProviderProductionSet provides a Loader for use in production.
	ProviderProductionSet = wire.NewSet(ProvideLoader, ProvideMeetingPlatformLoader)
)
