package domain_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/jake-hansen/agora/api/domain"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		username := "test"
		password := "test"
		fname := "test"
		lname := "test"

		user := domain.NewUser(fname, lname, username, password)

		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, fname, user.Firstname)
		assert.Equal(t, lname, user.Lastname)
	})
}
