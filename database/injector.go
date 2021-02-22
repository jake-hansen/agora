// +build wireinject

package database

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/log"
	"github.com/spf13/viper"
)

func Build(v *viper.Viper, log *log.Log) (*Manager, func(), error) {
	panic(wire.Build(ProviderProductionSet))
}

func BuildTest(cfg Config) (*MockManager, func(), error) {
	panic(wire.Build(ProviderTestSet))
}
