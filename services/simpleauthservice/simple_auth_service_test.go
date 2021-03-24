package simpleauthservice_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func ProvideTest() (*simpleauthservice.SimpleAuthService, *jwtservicemock.Service, *userservicemock.UserService) {
	jwtServiceMock := jwtservicemock.Provide()
	userServiceMock := userservicemock.Provide()
	authService := simpleauthservice.Provide(jwtServiceMock, userServiceMock)
	return authService, jwtServiceMock, userServiceMock
}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		as, js, us := ProvideTest()
		js.On("GenerateToken", mock.Anything).Return("test-token", nil)
		us.On("Validate", mock.Anything).Return(&domain.User{}, nil)
		token, err := as.Authenticate(testAuth)

		assert.NoError(t, err)

		js.On("ValidateToken", mock.Anything).Return(&jwt.Token{}, &jwtservice.Claims{}, nil)
		valid, err := as.IsAuthenticated(*token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})

	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.Token{Auth: "invalid"}
		as, ds, _ := ProvideTest()

		ds.On("ValidateToken", mock.Anything).Return(&jwt.Token{}, &jwtservice.Claims{}, errors.New("invalid token"))
		valid, err := as.IsAuthenticated(invalidToken)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

// TODO: cannot test until database is integrated
func TestSimpleAuthService_Authenticate(t *testing.T) {

}

func TestSimpleAuthService_Deauthenticate(t *testing.T) {
	invalidToken := domain.Token{Auth: "invalid"}
	as, _, _ := ProvideTest()
	err := as.Deauthenticate(invalidToken)

	assert.NoError(t, err)
}
