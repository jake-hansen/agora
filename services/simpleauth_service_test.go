package services_test

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var dur, _ = time.ParseDuration("5m")
var jwtService = services.ProvideJWTService(&testConfig)
var testAuth = domain.Auth{Credentials: &domain.User{
	Username:  "test",
	Password:  "test",
}}

func TestSimpleAuthService_IsAuthenticated(t *testing.T) {
	t.Run("test-valid-token", func(t *testing.T) {
		authService := services.ProvideSimpleAuthService(jwtService)
		token, err := authService.Authenticate(testAuth)

		assert.NoError(t, err)

		valid, err := authService.IsAuthenticated(*token)
		assert.NoError(t, err)
		assert.True(t, valid)
	})
	
	t.Run("test-invalid-token", func(t *testing.T) {
		invalidToken := domain.Token{Value: "invalid"}
		authService := services.ProvideSimpleAuthService(jwtService)

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
	authService := services.ProvideSimpleAuthService(jwtService)
	err := authService.Deauthenticate(invalidToken)

	assert.NoError(t, err)
}
