package jwtservicemock

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/stretchr/testify/mock"
	"time"
)

// Service is a mock of a JWT Service.
type Service struct {
	mock.Mock
}

// GenerateToken mocks JWTService's GenerateToken function.
func (j *Service) GenerateToken(user domain.User) (*domain.Token, error) {
	args := j.Mock.Called(user)
	return args.Get(0).(*domain.Token), args.Error(1)
}

// ValidateToken mocks JWTService's ValidateToken function.
func (j *Service) ValidateAuthToken(token string) (*jwt.Token, *jwtservice.Claims, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(*jwt.Token), args.Get(1).(*jwtservice.Claims), args.Error(2)
}

func (j *Service) GenerateRefreshToken(user domain.User, authToken domain.Token, expiry *time.Time) (*domain.Token, error) {
	args := j.Mock.Called(user, authToken, expiry)
	return args.Get(0).(*domain.Token), args.Error(1)
}

func (j *Service) ValidateRefreshToken(tokenSet domain.TokenSet) (*jwt.Token, *jwtservice.RefreshClaims, error) {
	args := j.Mock.Called(tokenSet)
	return args.Get(0).(*jwt.Token), args.Get(1).(*jwtservice.RefreshClaims), args.Error(2)
}

