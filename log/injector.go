// +build wireinject

package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func Build(v *viper.Viper) (*Log, func(), error) {
	panic(wire.Build(ProviderProductionSet))
}
