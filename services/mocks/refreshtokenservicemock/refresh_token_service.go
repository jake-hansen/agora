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

func (r *RefreshTokenService) ReplaceRefreshToken(token domain.RefreshToken) error {
	args := r.Called(token)
	return args.Error(0)
}

func (r *RefreshTokenService) GetRefreshTokenByParentTokenHash(token domain.RefreshToken) (*domain.RefreshToken, error) {
	args := r.Called(token)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}

func (r *RefreshTokenService) RevokeLatestRefreshTokenByNonce(token domain.RefreshToken) error {
	args := r.Called(token)
	return args.Error(0)
}

