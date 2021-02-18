package meetingplatformservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(repository domain.MeetingPlatformRepository) *MeetingPlatformService {
	return &MeetingPlatformService{repo: repository}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformService), new(*MeetingPlatformService)))
)