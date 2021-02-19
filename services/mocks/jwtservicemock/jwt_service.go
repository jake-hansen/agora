package jwtservicemock

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/stretchr/testify/mock"
)

// Service is a mock of a JWT Service.
type Service struct {
	mock.Mock
}

// GenerateToken mocks JWTService's GenerateToken function.
func (j *Service) GenerateToken(user domain.User) (string, error) {
	args := j.Mock.Called(user)
	return args.String(0), args.Error(1)
}

// ValidateToken mocks JWTService's ValidateToken function.
func (j *Service) ValidateToken(token string) (*jwt.Token, *jwtservice.Claims, error) {
	args := j.Mock.Called(token)
	return args.Get(0).(*jwt.Token), args.Get(1).(*jwtservice.Claims), args.Error(2)
}
