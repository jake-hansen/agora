// +build wireinject

package loader

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/services/meetingplatforms"
)

func Build() (*Loader, func(), error) {
	panic(wire.Build(ProviderProductionSet, meetingplatformrepo.ProviderSet, meetingplatforms.ProviderSet, database.ProviderProductionSet))
}
