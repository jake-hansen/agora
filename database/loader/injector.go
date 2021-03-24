// +build wireinject

package loader

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/platforms"
	"github.com/spf13/viper"
)

// Build returns a Loader configured with the provided Manager and Viper.
func Build(db *database.Manager, v *viper.Viper) (*Loader, error) {
	panic(wire.Build(ProviderProductionSet, meetingplatformrepo.ProviderSet, platforms.ProviderSet))
}
