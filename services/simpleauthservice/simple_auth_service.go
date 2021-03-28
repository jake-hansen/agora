package simpleauthservice

import (
	"errors"
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
)

// SimpleAuthService is an AuthenticationService which authenticates credentials based on a username
// and password combination. SimpleAuthService uses a JWT as a token which is not stored or persisted
// in any way. It is up to the consumer to reauthenticate upon JWT expiry to ensure continued access.
type SimpleAuthService struct {
	jwtService  jwtservice.JWTService
	userService domain.UserService
	refreshTokenService domain.RefreshTokenService
}

// IsAuthenticated determines whether the given Auth is authenticated. An Auth struct is considered authenticated
// if the contained JWT is valid.
func (s *SimpleAuthService) IsAuthenticated(token domain.AuthToken) (bool, error) {
	_, _, err := s.jwtService.ValidateAuthToken(token.Value)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Authenticate attempts to authenticate the given Auth. If authenticated, returns a AuthToken. Otherwise,
// an error is returned.
func (s *SimpleAuthService) Authenticate(auth domain.Auth) (*domain.TokenSet, error) {
	// Validate credentials with database
	if u, err := s.userService.Validate(auth.Credentials); err != nil {
		return nil, errors.New("username or password is not correct")
	} else {
		authToken, err2 := s.jwtService.GenerateAuthToken(*u)
		if err2 != nil {
			return nil, err
		}

		refreshToken, err2 := s.jwtService.GenerateRefreshToken(*u, *authToken, nil, nil)
		if err2 != nil {
			return nil, err2
		}

		_, err = s.refreshTokenService.SaveNewRefreshToken(*refreshToken)
		if err != nil {
			return nil, err
		}

		return &domain.TokenSet{
			Auth:    *authToken,
			Refresh: *refreshToken,
		}, nil
	}
}

func (s *SimpleAuthService) RefreshToken(token domain.RefreshToken) (*domain.TokenSet, error) {
	_, claims, err := s.jwtService.ValidateRefreshToken(token.Value)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	newAuthToken, err := s.jwtService.GenerateAuthToken(*user)
	if err != nil {
		return nil, err
	}

	expiry := time.Unix(claims.ExpiresAt, 0)
	newRefreshToken, err := s.jwtService.GenerateRefreshToken(*user, *newAuthToken, &token.Value, &expiry)
	if err != nil {
		return nil, err
	}

	newTokenSet := &domain.TokenSet{
		Auth:    *newAuthToken,
		Refresh: *newRefreshToken,
	}

	return newTokenSet, nil
}

// Deauthenticate is not implemented since JWTs are not persisted in a database.
func (s *SimpleAuthService) Deauthenticate(token domain.AuthToken) error {
	return nil
}

// GetUserFromAuthToken retrieves the User that the provided Token belongs to.
func (s *SimpleAuthService) GetUserFromAuthToken(token domain.AuthToken) (*domain.User, error) {
	_, claims, err := s.jwtService.ValidateAuthToken(token.Value)

	if err == nil && claims != nil {
		return s.userService.GetByID(claims.UserID)
	} else {
		return nil, err
	}
}
