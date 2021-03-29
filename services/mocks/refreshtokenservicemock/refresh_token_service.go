package refreshtokenservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

type RefreshTokenService struct {
	mock.Mock
}

func (r *RefreshTokenService) SaveNewRefreshToken(token domain.RefreshToken) (uint, error) {
	args := r.Called(token)
	return uint(args.Int(0)), args.Error(1)
}

func (r *RefreshTokenService) GetRefreshTokenByHash(hash string) (*domain.RefreshToken, error) {
	args := r.Called(hash)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}

func (r *RefreshTokenService) ReplaceRefreshToken(token domain.RefreshToken) error {
	args := r.Called(token)
	return args.Error(0)
}

func (r *RefreshTokenService) GetRefreshTokenByParentTokenHash(hash string, nonce string) (*domain.RefreshToken, error) {
	args := r.Called(hash, nonce)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}

func (r *RefreshTokenService) RevokeRefreshTokenByNonce(nonce string) error {
	args := r.Called(nonce)
	return args.Error(0)
}

