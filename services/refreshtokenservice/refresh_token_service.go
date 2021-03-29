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

func (r *RefreshTokenService) GetRefreshTokenByParentTokenHash(hash string, nonce string) (*domain.RefreshToken, error) {
	return r.repo.GetByParentTokenHash(hash, r.sha256(nonce))
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

func (r *RefreshTokenService) RevokeRefreshTokenByNonce(nonce string) error {
	foundToken, err := r.repo.GetByTokenNonceHash(r.sha256(nonce))
	if err != nil {
		return err
	}

	foundToken.Revoked = true

	return r.repo.Update(foundToken)
}

func (r *RefreshTokenService) sha256(v string) string {
	hasher := sha256.New()
	hasher.Write([]byte(v))
	return hex.EncodeToString(hasher.Sum(nil))
}
