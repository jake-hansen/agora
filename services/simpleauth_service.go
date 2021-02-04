package services

import (
	"errors"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwt"
)

// SimpleAuthService is an AuthenticationService which authenticates credentials based on a username
// and password combination. SimpleAuthService uses a JWT as a token which is not stored or persisted
// in any way. It is up to the consumer to reauthenticate upon JWT expiry to ensure continued access.
type SimpleAuthService struct {
	tokenService *jwt.Service
	userService	 domain.UserService
}

// ProvideSimpleAuthService returns a new SimpleAuthService which uses the given JWTService for generating and validating
// JWTs.
func ProvideSimpleAuthService(tokenService *jwt.Service, userService domain.UserService) *SimpleAuthService {
	return &SimpleAuthService{tokenService: tokenService, userService: userService}
}

// IsAuthenticated determines whether the given Auth is authenticated. An Auth struct is considered authenticated
// if the contained JWT is valid.
func (s *SimpleAuthService) IsAuthenticated(token domain.Token) (bool, error) {
	_, err := s.tokenService.ValidateToken(token.Value)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Authenticate attempts to authenticate the given Auth. If authenticated, returns a JWT. Otherwise,
// an error is returned.
// TODO: implement database to validate username/password
func (s *SimpleAuthService) Authenticate(auth domain.Auth) (*domain.Token, error) {
	// Validate credentials with database

	if auth.Credentials.Username == "test" && auth.Credentials.Password == "test" {
		// Generate JWT
		token, err := s.tokenService.GenerateToken(*auth.Credentials)
		if err != nil {
			return nil, err
		}
		return &domain.Token{Value: token}, nil
	}
	return nil, errors.New("username or password is not correct")
}

// Deauthenticate is not implemented since JWTs are not persisted in a database.
func (s *SimpleAuthService) Deauthenticate(token domain.Token) error {
	return nil
}

var (
	SimpleAuthServiceSet = wire.NewSet(ProvideSimpleAuthService, jwt.JWTServiceSet)
)

