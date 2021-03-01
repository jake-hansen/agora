package meetingplatformrepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a MeetingPlatformRepo that is configured by the given Manager.
func Provide(manager *database.Manager) *MeetingPlatformRepo {
	return &MeetingPlatformRepo{DB: manager.DB}
}

var (
	// ProviderSet provides a MeetingPlatformRepo for use in production.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.MeetingPlatformRepository), new(*MeetingPlatformRepo)))
)
