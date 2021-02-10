package jwtservice_test

import (
	"testing"
	"time"

	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/jwtservice"
	"github.com/stretchr/testify/assert"
)

var testConfig = jwtservice.Config{
	Issuer:     "agora-test",
	SigningKey: "testkey",
	Duration:   300000000000,
}

var testUser domain.User = domain.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  domain.NewPassword("password123"),
}

func TestJWTService_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		service, err := jwtservice.BuildTest(testConfig)
		assert.NoError(t, err)
		_, err = service.GenerateToken(testUser)
		assert.NoError(t, err)
	})
}

func TestJWTService_ValidateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		service, err := jwtservice.BuildTest(testConfig)
		assert.NoError(t, err)
		token, err := service.GenerateToken(testUser)
		assert.NoError(t, err)

		parsedToken, err := service.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, token, parsedToken.Raw)
	})

	t.Run("fail-expired", func(t *testing.T) {
		c := testConfig
		c.Duration = 1
		service, err := jwtservice.BuildTest(c)
		assert.NoError(t, err)
		token, err := service.GenerateToken(testUser)

		assert.NoError(t, err)

		time.Sleep(1000000000)
		_, err = service.ValidateToken(token)
		assert.Error(t, err)
	})
}
