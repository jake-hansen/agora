package simpleinviteservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

func Provide(inviteRepo domain.InviteRepository) *SimpleInviteService {
	return &SimpleInviteService{
		inviteRepo: inviteRepo,
	}
}

var (
	// ProviderProductionSet provides a new SimpleInviteService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.InviteService), new(*SimpleInviteService)))
)
