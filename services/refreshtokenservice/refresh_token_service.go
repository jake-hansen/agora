package refreshtokenservice

import (
	"github.com/jake-hansen/agora/domain"
)

// RefreshTokenService manages operations on RefreshTokens.
type RefreshTokenService struct {
	repo domain.RefreshTokenRepository
}

// SaveNewRefreshToken saves a new RefreshToken to the repository.
func (r *RefreshTokenService) SaveNewRefreshToken(token domain.RefreshToken) (uint, error) {
	return r.repo.Create(token)
}

// ReplaceRefreshToken deletes the existing RefreshToken in the repository that has the same
// nonce hash as the provided RefreshToken and then stores the provided RefreshToken
// in the repository.
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

// RevokeLatestRefreshTokenByNonce revokes the first found RefreshToken in the repository that
// contains the same nonce as the provided RefreshToken.
func (r *RefreshTokenService) RevokeLatestRefreshTokenByNonce(token domain.RefreshToken) error {
	foundToken, err := r.repo.GetByTokenNonceHash(token.TokenNonceHash)
	if err != nil {
		return err
	}

	foundToken.Revoked = true

	return r.repo.Update(foundToken)
}

// GetLatestTokenInSession gets the first found RefreshToken in the repository that has the same nonce
// hash as the provided RefreshToken.
func (r *RefreshTokenService) GetLatestTokenInSession(token domain.RefreshToken) (*domain.RefreshToken, error) {
	return r.repo.GetByTokenNonceHash(token.TokenNonceHash)
}
