package meetingplatformservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(repository domain.MeetingPlatformRepository, configuredPlatforms domain.ConfiguredPlatforms) *MeetingPlatformService {
	return New(repository, configuredPlatforms)
}
var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformService), new(*MeetingPlatformService)))
)