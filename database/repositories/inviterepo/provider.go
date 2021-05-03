package inviterepo

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
)

// Provide returns an InviteRepo configured with the provided Manager.
func Provide(manager *database.Manager) *InviteRepo {
	return &InviteRepo{DB: manager.DB}
}

var (
	// ProviderSet provides a MeetingPlatformRepo for use in production.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.InviteRepository), new(*InviteRepo)))
)
