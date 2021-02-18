package meetingplatformrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

func Provide(manager *database.Manager) *MeetingPlatformRepo {
	return &MeetingPlatformRepo{DB: manager.DB}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformRepository), new(*MeetingPlatformRepo)))
)
