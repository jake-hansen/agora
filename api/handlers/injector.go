// +build wireinject

package handlers

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/router/handlers"
	"github.com/spf13/viper"
)

func Build(db *database.Manager, v *viper.Viper) (*[]handlers.Handler, func(), error) {
	panic(wire.Build(ProviderProductionSet))
}
