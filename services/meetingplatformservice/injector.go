// +build wireinject

package meetingplatformservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func BuildTest(repo domain.MeetingPlatformRepository) *MeetingProviderService {
	panic(wire.Build(Provide))
}