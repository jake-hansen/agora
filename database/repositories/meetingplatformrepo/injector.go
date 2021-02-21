// +build wireinject

package meetingplatformrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
)

func Build(manager *database.Manager) (*MeetingPlatformRepo, func(), error)  {
	panic(wire.Build(Provide))
}