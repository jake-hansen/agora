package userservice_test

import (
	"errors"
	"testing"

	"github.com/jake-hansen/agora/database/repositories/mocks/userrepomock"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/userservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUserPassword = "Password123"

var passwordObj = domain.NewPassword("Password123")

var mockUser = domain.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  passwordObj,
}

var mockUserHash = "$2a$10$PdjlGYhMGonCrjKNquZmzeMQY0M4vlxsCjtQysCOOSzxcfpTW5JAe"

func TestUserService_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Create", mock.AnythingOfType("*domain.User")).Return(1, nil)

		uService := userservice.BuildTest(r)

		id, err := uService.Register(&mockUser)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
	})

	t.Run("failure", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Create", mock.AnythingOfType("*domain.User")).Return(0, errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		id, err := uService.Register(&mockUser)

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})
}

func TestUserService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsers := []*domain.User{&mockUser, &mockUser}

		r := userrepomock.Build()
		r.On("GetAll").Return(mockUsers, nil)

		uService := userservice.BuildTest(r)

		users, err := uService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, mockUsers, users)
	})

	t.Run("failure", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetAll").Return([]*domain.User{}, errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		users, err := uService.GetAll()

		assert.Error(t, err)
		assert.Empty(t, users)
	})
}

func TestUserService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&mockUser, nil)

		uService := userservice.BuildTest(r)

		user, err := uService.GetByID(0)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, *user)
	})

	t.Run("failure", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&domain.User{}, errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		user, err := uService.GetByID(0)

		assert.Error(t, err)
		assert.Empty(t, user)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)

		uService := userservice.BuildTest(r)

		err := uService.Update(&mockUser)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Update", mock.AnythingOfType("*domain.User")).Return(errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		err := uService.Update(&mockUser)

		assert.Error(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Delete", mock.AnythingOfType("uint")).Return(nil)

		uService := userservice.BuildTest(r)

		err := uService.Delete(0)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("Delete", mock.AnythingOfType("uint")).Return(errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		err := uService.Delete(0)

		assert.Error(t, err)
	})
}

func TestUserService_Validate(t *testing.T) {
	returnUser := mockUser
	returnUser.Password.Hash = []byte(mockUserHash)

	t.Run("success", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&returnUser, nil)

		uService := userservice.BuildTest(r)

		creds := &domain.Credentials{
			Username: mockUser.Username,
			Password: mockUserPassword,
		}

		_, err := uService.Validate(creds)

		assert.NoError(t, err)
	})

	t.Run("bad-password", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&returnUser, nil)

		uService := userservice.BuildTest(r)

		creds := &domain.Credentials{
			Username: mockUser.Username,
			Password: "wrong-password",
		}

		_, err := uService.Validate(creds)

		assert.Error(t, err)
	})

	t.Run("retrieval-error", func(t *testing.T) {
		r := userrepomock.Build()
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("unknown error"))

		uService := userservice.BuildTest(r)

		creds := &domain.Credentials{
			Username: mockUser.Username,
			Password: "wrong-password",
		}

		_, err := uService.Validate(creds)

		assert.Error(t, err)
	})
}
