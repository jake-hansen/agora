package simpleauthservice

import (
	"errors"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
)

// SimpleAuthService is an AuthenticationService which authenticates credentials based on a username
// and password combination. SimpleAuthService issues AuthTokens and RefreshTokens which can be used
// in combination to persist a session.
type SimpleAuthService struct {
	jwtService          jwtservice.JWTService
	userService         domain.UserService
	refreshTokenService domain.RefreshTokenService
}

// IsAuthenticated determines whether the given TokenValue is a valid AuthToken.
func (s *SimpleAuthService) IsAuthenticated(token domain.TokenValue) (bool, error) {
	_, err := s.jwtService.ValidateAuthToken(token)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Authenticate attempts to authenticate the given Auth. If authenticated, returns a TokenSet. Otherwise,
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

// RefreshToken attempts to refresh a user's auth token provided that the given TokenValue
// is a valid RefreshToken.
func (s *SimpleAuthService) RefreshToken(token domain.TokenValue) (*domain.TokenSet, error) {
	parsedToken, err := s.jwtService.ValidateRefreshToken(token)
	if err != nil {
		return nil, err
	}
	latestToken, err := s.refreshTokenService.GetLatestTokenInSession(parsedToken)
	if err != nil {
		return nil, err
	} else {
		if string(latestToken.Value) != parsedToken.ParentTokenHash && latestToken.ParentTokenHash != "" {
			_ = s.refreshTokenService.RevokeLatestRefreshTokenByNonce(parsedToken)
			return nil, NewRefreshTokenReuseError()
		} else if latestToken.Revoked {
			return nil, NewRefreshTokenRevokedError()
		}
	}

	user, err := s.userService.GetByID(parsedToken.UserID)
	if err != nil {
		return nil, err
	}

	var newRefreshToken *domain.RefreshToken

	newAuthToken, err := s.jwtService.GenerateAuthToken(*user)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err = s.jwtService.GenerateRefreshToken(*user, *newAuthToken, &parsedToken)
	if err != nil {
		return nil, err
	}

	if parsedToken.ParentTokenHash != "" {
		err = s.refreshTokenService.ReplaceRefreshToken(parsedToken)
	}

	newTokenSet := &domain.TokenSet{
		Auth:    *newAuthToken,
		Refresh: *newRefreshToken,
	}

	return newTokenSet, nil
}

// Deauthenticate revokes the provided RefreshToken.
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
