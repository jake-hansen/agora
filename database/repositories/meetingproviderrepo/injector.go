// +build wireinject

package meetingproviderrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
)

func Build(manager *database.Manager) (*MeetingProviderRepo, func(), error)  {
	panic(wire.Build(Provide))
}