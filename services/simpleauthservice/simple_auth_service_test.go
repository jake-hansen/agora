package simpleauthservice_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jake-hansen/agora/services/mocks/refreshtokenservicemock"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/mocks/jwtservicemock"
	"github.com/jake-hansen/agora/services/mocks/userservicemock"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dur, _ = time.ParseDuration("5m")
var testconfig = jwtservice.Config{
	Issuer:     "test",
	SigningKey: "test",
	Duration:   dur,
}
var testAuth = domain.Auth{Credentials: &domain.Credentials{
	Username: "test",
	Password: "test",
}}

func ProvideTest() (*simpleauthservice.SimpleAuthService, *jwtservicemock.Service, *userservicemock.UserService, *refreshtokenservicemock.RefreshTokenService) {
	jwtServiceMock := jwtservicemock.Provide()
	userServiceMock := userservicemock.Provide()
	refreshTokenServiceMock := refreshtokenservicemock.Provide()
	authService := simpleauthservice.Provide(jwtServiceMock, userServiceMock, refreshTokenServiceMock)
	return authService, jwtServiceMock, userServiceMock, refreshTokenServiceMock
}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		as, js, us, rs := ProvideTest()
		testAuthToken := domain.AuthToken{Value: "test-token"}
		testRefreshToken := domain.RefreshToken{Value: "test-refresh-token"}
		js.On("GenerateAuthToken", mock.Anything).Return(&testAuthToken, nil)
		js.On("GenerateRefreshToken", mock.Anything, mock.Anything, mock.Anything).Return(&testRefreshToken, nil)
		us.On("Validate", mock.Anything).Return(&domain.User{}, nil)
		rs.On("SaveNewRefreshToken", mock.AnythingOfType("domain.RefreshToken")).Return(0, nil)
		token, err := as.Authenticate(testAuth)

		assert.NoError(t, err)

		js.On("ValidateAuthToken", mock.Anything).Return(domain.AuthToken{}, nil)
		valid, err := as.IsAuthenticated(token.Auth.Value)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.AuthToken{Value: "invalid"}
		as, ds, _, _ := ProvideTest()

		ds.On("ValidateAuthToken", mock.Anything).Return(domain.AuthToken{}, errors.New("invalid token"))
		valid, err := as.IsAuthenticated(invalidToken.Value)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

// TODO: cannot test until database is integrated
func TestSimpleAuthService_Authenticate(t *testing.T) {

}

func TestSimpleAuthService_Deauthenticate(t *testing.T) {
	invalidToken := domain.RefreshToken{Value: "invalid"}
	as, js, _, _ := ProvideTest()
	js.On("ValidateRefreshToken", mock.AnythingOfType("domain.TokenValue")).Return(domain.RefreshToken{}, errors.New("test error"))
	err := as.Deauthenticate(invalidToken.Value)

	assert.Error(t, err)
}
