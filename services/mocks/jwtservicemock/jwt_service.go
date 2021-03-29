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
func (j *Service) GenerateAuthToken(user domain.User) (*domain.AuthToken, error) {
	args := j.Mock.Called(user)
	return args.Get(0).(*domain.AuthToken), args.Error(1)
}

// ValidateToken mocks JWTService's ValidateToken function.
func (j *Service) ValidateAuthToken(token string) (*jwt.Token, *jwtservice.AuthClaims, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(*jwt.Token), args.Get(1).(*jwtservice.AuthClaims), args.Error(2)
}

func (j *Service) GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshTokenValue, expiry *time.Time) (*domain.RefreshToken, error) {
	args := j.Mock.Called(user, authToken, parentToken, expiry)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}

func (j *Service) ValidateRefreshToken(token domain.RefreshTokenValue) (*jwt.Token, *jwtservice.RefreshClaims, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(*jwt.Token), args.Get(1).(*jwtservice.RefreshClaims), args.Error(2)
}

