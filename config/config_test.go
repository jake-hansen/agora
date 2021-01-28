package config_test

import (
	"github.com/jake-hansen/agora/config"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	t.Run("success-dev", func(t *testing.T) {
		config.Init("dev")
	})
	t.Run("success-prod", func(t *testing.T) {
		config.Init("dev")
	})
	t.Run("success-test", func(t *testing.T) {
		config.Init("dev")
	})
	t.Run("fail-not-found", func(t *testing.T) {
		assert.Panics(t, func() { config.Init("non-existent") })
	})
}

func TestGetConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		config.Init("dev")
		config := config.GetConfig()

		assert.True(t, strings.Contains(config.ConfigFileUsed(), "dev.yaml"))
	})
}
