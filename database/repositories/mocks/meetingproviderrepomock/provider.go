package meetingproviderrepomock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

func Provide() *MeetingProviderRepository {
	return &MeetingProviderRepository{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingProviderRepository), new(*MeetingProviderRepository)))
)