// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func Build() (*viper.Viper, func(), error) {
	panic(wire.Build(Provide))
}
