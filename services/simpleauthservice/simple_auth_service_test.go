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
var testconfig = jwtservice.Config{
	Issuer:     "test",
	SigningKey: "test",
	Duration:   dur,
}
var testAuth = domain.Auth{Credentials: &domain.User{
	Username:  "test",
	Password:  "test",
}}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		authService, err := simpleauthservice.BuildTest(testconfig)
		assert.NoError(t, err)
		token, err := authService.Authenticate(testAuth)

		assert.NoError(t, err)

		valid, err := authService.IsAuthenticated(*token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})
	
	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.Token{Value: "invalid"}
		authService, err := simpleauthservice.BuildTest(testconfig)
		assert.NoError(t, err)

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
	authService, err := simpleauthservice.BuildTest(testconfig)
	assert.NoError(t, err)
	err = authService.Deauthenticate(invalidToken)

	assert.NoError(t, err)
}
