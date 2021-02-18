// +build wireinject

package meetingplatformrepomock

import "github.com/google/wire"

func Build() *MeetingPlatformRepository {
	panic(wire.Build(ProviderSet))
}
