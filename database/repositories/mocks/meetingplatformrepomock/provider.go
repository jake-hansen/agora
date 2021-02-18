package meetingplatformrepomock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

func Provide() *MeetingPlatformRepository {
	return &MeetingPlatformRepository{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformRepository), new(*MeetingPlatformRepository)))
)