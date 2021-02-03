package services_test

import (
	"github.com/jake-hansen/agora/api/dto"
	"github.com/jake-hansen/agora/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testUser dto.User = dto.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  "password123",
}

func TestJWTService_GenerateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dur, err := time.ParseDuration("5m")
		if err != nil {
			panic(err)
		}
		jwt := services.NewJWTService("agora-test", "testkey", dur)
		_, err = jwt.GenerateToken(testUser)

		assert.NoError(t, err)
	})
}

func TestJWTService_ValidateToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dur, err := time.ParseDuration("5m")
		if err != nil {
			panic(err)
		}
		jwt := services.NewJWTService("agora-test", "testkey", dur)
		token, err := jwt.GenerateToken(testUser)

		assert.NoError(t, err)

		parsedToken, err := jwt.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, token, parsedToken.Raw)
	})

	t.Run("fail-expired", func(t *testing.T) {
		dur, err := time.ParseDuration("1µs")
		if err != nil {
			panic(err)
		}
		jwt := services.NewJWTService("agora-test", "testkey", dur)
		token, err := jwt.GenerateToken(testUser)

		assert.NoError(t, err)

		time.Sleep(1 * time.Second)
		_, err = jwt.ValidateToken(token)
		assert.Error(t, err)
	})
}
