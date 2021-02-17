// +build wireinject

package meetingproviderrepomock

import "github.com/google/wire"

func Build() *MeetingProviderRepository {
	panic(wire.Build(ProviderSet))
}
