package refreshtokenservice

import (
	"github.com/jake-hansen/agora/domain"
)

type RefreshTokenService struct {
	repo domain.RefreshTokenRepository
}

func (r *RefreshTokenService) SaveNewRefreshToken(token domain.RefreshToken) (uint, error) {
	return r.repo.Create(token)
}

func (r *RefreshTokenService) ReplaceRefreshToken(token domain.RefreshToken) error {
	nonceFamily := token.TokenNonceHash

	foundToken, err := r.repo.GetByTokenNonceHash(nonceFamily)
	if err != nil {
		return err
	}

	err = r.repo.Delete(foundToken.ID)
	if err != nil {
		return err
	}

	_, err = r.repo.Create(token)
	return err
}

func (r *RefreshTokenService) RevokeLatestRefreshTokenByNonce(token domain.RefreshToken) error {
	foundToken, err := r.repo.GetByTokenNonceHash(token.TokenNonceHash)
	if err != nil {
		return err
	}

	foundToken.Revoked = true

	return r.repo.Update(foundToken)
}

func (r *RefreshTokenService) GetLatestTokenInSession(token domain.RefreshToken) (*domain.RefreshToken, error) {
	return r.repo.GetByTokenNonceHash(token.TokenNonceHash)
}
