package meetingplatformservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns a MeetingPlatformService configured with the provided MeetingPlatformRepository and ConfiguredPlatforms.
func Provide(repository domain.MeetingPlatformRepository, configuredPlatforms domain.ConfiguredPlatforms) *MeetingPlatformService {
	return New(repository, configuredPlatforms)
}

var (
	// ProviderSet provides a MeetingPlatformService.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformService), new(*MeetingPlatformService)))
)
