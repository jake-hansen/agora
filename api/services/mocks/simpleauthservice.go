package mocks

import (
	"github.com/jake-hansen/agora/api/domain"
	"github.com/stretchr/testify/mock"
)

type SimpleAuthService struct {
	mock.Mock
}

func (s *SimpleAuthService) IsAuthenticated(auth domain.Auth) (bool, error) {
	args := s.Called()
	return args.Bool(0), args.Error(1)
}

func (s *SimpleAuthService) Authenticate(auth domain.Auth) (interface{}, error) {
	args := s.Called()
	return args.Get(0), args.Error(1)
}

func (s *SimpleAuthService) Deauthenticate(auth domain.Auth) error {
	args := s.Called()
	return args.Error(0)
}

