//+build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/log"
	"github.com/jake-hansen/agora/router"
	"github.com/spf13/viper"
)

func Build(db *database.Manager, v *viper.Viper, log *log.Log) (*Server, error) {
	panic(wire.Build(ProviderProductionSet,
			         router.ProviderProductionSet,
			         handlers.ProviderProductionSet,
			         middleware.ProviderProductionSet))
}
