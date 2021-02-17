package meetingproviderservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(repository domain.MeetingProviderRepository) *MeetingProviderService {
	return &MeetingProviderService{repo: repository}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingProviderService), new(*MeetingProviderService)))
)