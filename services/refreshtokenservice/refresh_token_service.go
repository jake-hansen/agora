package refreshtokenservice

import "github.com/jake-hansen/agora/domain"

type RefreshTokenService struct {
	repo domain.RefreshTokenRepository
}

func (r *RefreshTokenService) SaveNewRefreshToken(token domain.RefreshToken) (uint, error) {
	return r.repo.Create(token)
}

