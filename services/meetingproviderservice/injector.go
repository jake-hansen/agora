// +build wireinject

package meetingproviderservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func BuildTest(repo domain.MeetingProviderRepository) *MeetingProviderService {
	panic(wire.Build(Provide))
}