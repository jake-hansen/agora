package refreshtokenservice

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/jake-hansen/agora/domain"
)

type RefreshTokenService struct {
	repo domain.RefreshTokenRepository
}

func (r *RefreshTokenService) SaveNewRefreshToken(token domain.RefreshToken) (uint, error) {
	return r.repo.Create(token)
}

func (r *RefreshTokenService) GetRefreshTokenByHash(hash string) (*domain.RefreshToken, error) {
	return r.repo.GetByTokenHash(hash)
}

func (r *RefreshTokenService) GetRefreshTokenByParentTokenHash(hash string) (*domain.RefreshToken, error) {
	return r.repo.GetByParentTokenHash(hash)
}

func (r *RefreshTokenService) GetRefreshTokenByTokenNonceHash(nonceHash string) (*domain.RefreshToken, error) {
	return r.repo.GetByTokenNonceHash(nonceHash)
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

func (r *RefreshTokenService) RevokeRefreshTokenByNonce(nonceHash string) error {
	hasher := sha256.New()
	hasher.Write([]byte(nonceHash))
	hash := hex.EncodeToString(hasher.Sum(nil))

	foundToken, err := r.repo.GetByTokenNonceHash(hash)
	if err != nil {
		return err
	}

	foundToken.Revoked = true

	return r.repo.Update(foundToken)
}

