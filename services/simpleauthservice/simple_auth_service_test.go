package simpleauthservice_test

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/mocks/jwtservicemock"
	"github.com/jake-hansen/agora/services/mocks/userservicemock"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var dur, _ = time.ParseDuration("5m")
var testconfig = jwtservice.Config{
	Issuer:     "test",
	SigningKey: "test",
	Duration:   dur,
}
var testAuth = domain.Auth{Credentials: &domain.User{
	Username:  "test",
	Password:  "test",
}}

func buildTest(t *testing.T) (*simpleauthservice.SimpleAuthService, *jwtservicemock.Service, *userservicemock.UserService) {
	jwtServiceMock := jwtservicemock.Build()
	userServiceMock := userservicemock.Build()
	authService, err := simpleauthservice.BuildTest(jwtServiceMock, userServiceMock)
	assert.NoError(t, err)
	return authService, jwtServiceMock, userServiceMock
}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		as, js, _ := buildTest(t)
		js.On("GenerateToken", mock.Anything).Return("test-token", nil)
		token, err := as.Authenticate(testAuth)

		assert.NoError(t, err)

		js.On("ValidateToken", mock.Anything).Return(&jwt.Token{}, nil)
		valid, err := as.IsAuthenticated(*token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})
	
	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.Token{Value: "invalid"}
		as, ds, _ := buildTest(t)

		ds.On("ValidateToken", mock.Anything).Return(&jwt.Token{}, errors.New("invalid token"))
		valid, err := as.IsAuthenticated(invalidToken)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

// TODO: cannot test until database is integrated
func TestSimpleAuthService_Authenticate(t *testing.T) {

}

func TestSimpleAuthService_Deauthenticate(t *testing.T) {
	invalidToken := domain.Token{Value: "invalid"}
	as, _, _ := buildTest(t)
	err := as.Deauthenticate(invalidToken)

	assert.NoError(t, err)
}
