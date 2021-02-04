//+build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func BuildConfig() *viper.Viper {
	panic(wire.Build(ProvideViper))
}
