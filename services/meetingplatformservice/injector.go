// +build wireinject

package meetingplatformservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func BuildTest(repo domain.MeetingPlatformRepository, configuredPlatforms []*domain.MeetingPlatform) *MeetingPlatformService {
	panic(wire.Build(Provide))
}