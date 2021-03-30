package simpleauthservice

import (
	"errors"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"gorm.io/gorm"
)

// SimpleAuthService is an AuthenticationService which authenticates credentials based on a username
// and password combination. SimpleAuthService uses a JWT as a token which is not stored or persisted
// in any way. It is up to the consumer to reauthenticate upon JWT expiry to ensure continued access.
type SimpleAuthService struct {
	jwtService          jwtservice.JWTService
	userService         domain.UserService
	refreshTokenService domain.RefreshTokenService
}

// IsAuthenticated determines whether the given Auth is authenticated. An Auth struct is considered authenticated
// if the contained JWT is valid.
func (s *SimpleAuthService) IsAuthenticated(token domain.TokenValue) (bool, error) {
	_, err := s.jwtService.ValidateAuthToken(token)
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

		refreshToken, err2 := s.jwtService.GenerateRefreshToken(*u, *authToken, nil)
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

func (s *SimpleAuthService) RefreshToken(token domain.TokenValue) (*domain.TokenSet, error) {
	parsedToken, err := s.jwtService.ValidateRefreshToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetByID(parsedToken.UserID)
	if err != nil {
		return nil, err
	}

	var newRefreshToken *domain.RefreshToken

	foundToken, err := s.refreshTokenService.GetRefreshTokenByParentTokenHash(parsedToken)
	if err != nil {
		err = s.refreshTokenService.RevokeLatestRefreshTokenByNonce(parsedToken)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
			return nil, err
		}
		return nil, NewRefreshTokenReuseError()
	} else {
		if foundToken.Revoked == true {
			return nil, NewRefreshTokenRevokedError()
		}
	}

	newAuthToken, err := s.jwtService.GenerateAuthToken(*user)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err = s.jwtService.GenerateRefreshToken(*user, *newAuthToken, &parsedToken)
	if err != nil {
		return nil, err
	}

	err = s.refreshTokenService.ReplaceRefreshToken(*newRefreshToken)
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
func (s *SimpleAuthService) Deauthenticate(token domain.TokenValue) error {
	parsedToken, err := s.jwtService.ValidateRefreshToken(token)
	if err != nil {
		return err
	}

	return s.refreshTokenService.RevokeLatestRefreshTokenByNonce(parsedToken)
}

// GetUserFromAuthToken retrieves the User that the provided Token belongs to.
func (s *SimpleAuthService) GetUserFromAuthToken(token domain.TokenValue) (*domain.User, error) {
	parsedToken, err := s.jwtService.ValidateAuthToken(token)

	if err == nil {
		return s.userService.GetByID(parsedToken.JWTClaims.UserID)
	} else {
		return nil, err
	}
}
