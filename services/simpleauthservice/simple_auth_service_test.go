package simpleauthservice_test

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/jake-hansen/agora/services/simpleauthservice"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var dur, _ = time.ParseDuration("5m")
var jwtService = jwtservice.ProvideJWTService(&jwtservice.testConfig)
var testAuth = domain.Auth{Credentials: &domain.User{
	Username:  "test",
	Password:  "test",
}}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		authService := simpleauthservice.ProvideSimpleAuthService(jwtService)
		token, err := authService.Authenticate(testAuth)

		assert.NoError(t, err)

		valid, err := authService.IsAuthenticated(*token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})
	
	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.Token{Value: "invalid"}
		authService := simpleauthservice.ProvideSimpleAuthService(jwtService)

		valid, err := authService.IsAuthenticated(invalidToken)
		assert.Error(t, err)
		assert.False(t, valid)
	})
}

// TODO: cannot test until database is integrated
func TestSimpleAuthService_Authenticate(t *testing.T) {

}

func TestSimpleAuthService_Deauthenticate(t *testing.T) {
	invalidToken := domain.Token{Value: "invalid"}
	authService := simpleauthservice.ProvideSimpleAuthService(jwtService)
	err := authService.Deauthenticate(invalidToken)

	assert.NoError(t, err)
}
