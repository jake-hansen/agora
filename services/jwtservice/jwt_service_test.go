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

var testUser = domain.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  domain.NewPassword("password123"),
}

func TestJWTService_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		service := jwtservice.Provide(&testConfig)
		_, err := service.GenerateToken(testUser)
		assert.NoError(t, err)
	})
}

func TestJWTService_ValidateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		service := jwtservice.Provide(&testConfig)
		token, err := service.GenerateToken(testUser)
		assert.NoError(t, err)

		parsedToken, _, err := service.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, token, parsedToken.Raw)
	})

	t.Run("fail-expired", func(t *testing.T) {
		c := testConfig
		c.Duration = 1 * time.Microsecond
		service := jwtservice.Provide(&c)
		token, err := service.GenerateToken(testUser)

		assert.NoError(t, err)

		time.Sleep(1 * time.Second)
		_, _, err = service.ValidateToken(token)
		assert.Error(t, err)
	})
}
