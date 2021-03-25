package authservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// AuthService is a mock AuthService.
type AuthService struct {
	mock.Mock
}

// GetUser mocks AuthService's GetUser function.
func (s *AuthService) GetUser(token domain.Token) (*domain.User, error) {
	args := s.Called(token)
	return args.Get(0).(*domain.User), args.Error(1)
}

// IsAuthenticated mocks AuthService's IsAuthenticated function.
func (s *AuthService) IsAuthenticated(token domain.Token) (bool, error) {
	args := s.Called(token)
	return args.Bool(0), args.Error(1)
}

// Authenticate mocks AuthService's Authenticate function.
func (s *AuthService) Authenticate(auth domain.Auth) (*domain.TokenSet, error) {
	args := s.Called(auth)
	return args.Get(0).(*domain.TokenSet), args.Error(1)
}

func (s *AuthService) RefreshToken(tokens domain.TokenSet) (*domain.TokenSet, error) {
	args := s.Called(tokens)
	return args.Get(0).(*domain.TokenSet), args.Error(1)
}

// Deauthenticate mocks AuthService's Deauthenticate function.
func (s *AuthService) Deauthenticate(token domain.Token) error {
	args := s.Called(token)
	return args.Error(0)
}
