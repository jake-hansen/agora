package services

import (
	"errors"
	"github.com/jake-hansen/agora/api/domain"
)

// SimpleAuthService is an AuthenticationService which authenticates credentials based on a username
// and password combination. SimpleAuthService uses a JWT as a token.
type SimpleAuthService struct {
	tokenService *JWTService
}

// NewSimpleAuthService returns a new SimpleAuthService which uses the given JWTService for generating and validating
// JWTs.
func NewSimpleAuthService(tokenService JWTService) *SimpleAuthService {
	return &SimpleAuthService{tokenService: &tokenService}
}

// IsAuthenticated determines whether the given Auth is authenticated. An Auth struct is considered authenticated
// if the contained JWT is valid.
func (s *SimpleAuthService) IsAuthenticated(auth domain.Auth) (bool, error) {
	_, err := s.tokenService.ValidateToken(auth.Token)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Authenticate attempts to authenticate the given Auth. If authenticated, returns a JWT. Otherwise,
// an error is returned.
// TODO: implement database to validate username/password
func (s *SimpleAuthService) Authenticate(auth domain.Auth) (interface{}, error) {
	// Validate credentials with database

	if auth.Credentials.Username == "test" && auth.Credentials.Password == "test" {
		// Generate JWT
		token, err := s.tokenService.GenerateToken(*auth.Credentials)
		if err != nil {
			return nil, err
		}
		return domain.Auth{Token: token}, nil
	}
	return nil, errors.New("username or password is not correct")
}

// Deauthenticate is not implemented since JWTs are not persisted in a database.
func (s *SimpleAuthService) Deauthenticate(auth domain.Auth) error {
	panic("implement me")
}

