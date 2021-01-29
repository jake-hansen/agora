package services

import (
	"errors"
	"github.com/jake-hansen/agora/api/domain"
)

type SimpleAuth struct {
	TokenService JWTService
}

func (s *SimpleAuth) IsAuthenticated(auth domain.Auth) (bool, error) {
	_, err := s.TokenService.ValidateToken(auth.Token)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *SimpleAuth) Authenticate(auth domain.Auth) (interface{}, error) {
	// Validate credentials with database
	if auth.User.Username == "test" && auth.User.Password == "test" {
		// Generate JWT
		token, err := s.TokenService.GenerateToken(auth.User)
		if err != nil {
			return nil, err
		}
		return domain.Auth{Token: token}, nil
	}
	return nil, errors.New("username or password is not correct")
}

func (s *SimpleAuth) Deauthenticate(auth domain.Auth) error {
	panic("implement me")
}

