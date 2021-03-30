package jwtservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
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
func (j *Service) ValidateAuthToken(token domain.TokenValue) (domain.AuthToken, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(domain.AuthToken), args.Error(1)
}

func (j *Service) GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshToken) (*domain.RefreshToken, error) {
	args := j.Mock.Called(user, authToken, parentToken)
	return args.Get(0).(*domain.RefreshToken), args.Error(1)
}

func (j *Service) ValidateRefreshToken(token domain.TokenValue) (domain.RefreshToken, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(domain.RefreshToken), args.Error(1)
}
