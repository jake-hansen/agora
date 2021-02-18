// +build wireinject

package meetingplatforms

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/spf13/viper"
)

func Build(v *viper.Viper) []*domain.MeetingPlatform {
	panic(wire.Build(ProviderSet))
}