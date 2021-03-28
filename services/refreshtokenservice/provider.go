package refreshtokenservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new RefreshTokenService configured with the given RefreshTokenRepository.
func Provide(repository domain.RefreshTokenRepository) *RefreshTokenService {
	return &RefreshTokenService{repo: repository}
}

var (
	// ProviderProductionSet provides a new RefreshTokenService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.RefreshTokenService), new(*RefreshTokenService)))
)
