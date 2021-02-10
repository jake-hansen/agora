package authservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// AuthService is a mock AuthService.
type AuthService struct {
	mock.Mock
}

// IsAuthenticated mocks AuthService's IsAuthenticated function.
func (s *AuthService) IsAuthenticated(token domain.Token) (bool, error) {
	args := s.Called()
	return args.Bool(0), args.Error(1)
}

// Authenticate mocks AuthService's Authenticate function.
func (s *AuthService) Authenticate(auth domain.Auth) (*domain.Token, error) {
	args := s.Called()
	return args.Get(0).(*domain.Token), args.Error(1)
}

// Deauthenticate mocks AuthService's Deauthenticate function.
func (s *AuthService) Deauthenticate(token domain.Token) error {
	args := s.Called()
	return args.Error(0)
}
