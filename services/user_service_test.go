package services_test

import (
	"errors"
	repositorymocks "github.com/jake-hansen/agora/database/repositories/mocks"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var mockUser = domain.User{
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  "Password123",
}

var mockUserHash = "$2a$10$PdjlGYhMGonCrjKNquZmzeMQY0M4vlxsCjtQysCOOSzxcfpTW5JAe"

func TestUserService_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Create", mock.AnythingOfType("*domain.User")).Return(1, nil)

		uService := services.NewUserService(r)

		id, err := uService.Register(&mockUser)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
	})

	t.Run("failure", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Create", mock.AnythingOfType("*domain.User")).Return(0, errors.New("unknown error"))

		uService := services.NewUserService(r)

		id, err := uService.Register(&mockUser)

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})
}

func TestUserService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsers := []*domain.User{&mockUser, &mockUser}

		r := new(repositorymocks.UserRepository)
		r.On("GetAll").Return(mockUsers, nil)

		uService := services.NewUserService(r)

		users, err := uService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, mockUsers, users)
	})

	t.Run("failure", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("GetAll").Return([]*domain.User{}, errors.New("unknown error"))

		uService := services.NewUserService(r)

		users, err := uService.GetAll()

		assert.Error(t, err)
		assert.Empty(t, users)
	})
}

func TestUserService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&mockUser, nil)

		uService := services.NewUserService(r)

		user, err := uService.GetByID(0)

		assert.NoError(t, err)
		assert.Equal(t, mockUser, *user)
	})

	t.Run("failure", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&domain.User{}, errors.New("unknown error"))

		uService := services.NewUserService(r)

		user, err := uService.GetByID(0)

		assert.Error(t, err)
		assert.Empty(t, user)
	})
}

func TestUserService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)

		uService := services.NewUserService(r)

		err := uService.Update(&mockUser)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Update", mock.AnythingOfType("*domain.User")).Return(errors.New("unknown error"))

		uService := services.NewUserService(r)

		err := uService.Update(&mockUser)

		assert.Error(t, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Delete", mock.AnythingOfType("uint")).Return(nil)

		uService := services.NewUserService(r)

		err := uService.Delete(0)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("Delete", mock.AnythingOfType("uint")).Return(errors.New("unknown error"))

		uService := services.NewUserService(r)

		err := uService.Delete(0)

		assert.Error(t, err)
	})
}

func TestUserService_Validate(t *testing.T) {
	returnUser := mockUser
	returnUser.Password = mockUserHash

	t.Run("success", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		returnUser := mockUser
		returnUser.Password = mockUserHash
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&returnUser, nil)

		uService := services.NewUserService(r)

		vUser := &domain.User{
			Username: mockUser.Username,
			Password: mockUser.Password,
		}

		err := uService.Validate(vUser)

		assert.NoError(t, err)
	})

	t.Run("bad-password", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&returnUser, nil)

		uService := services.NewUserService(r)

		vUser := &domain.User{
			Username: mockUser.Username,
			Password: "wrong-password",
		}

		err := uService.Validate(vUser)

		assert.Error(t, err)
	})

	t.Run("retrieval-error", func(t *testing.T) {
		r := new(repositorymocks.UserRepository)
		r.On("GetByUsername", mock.AnythingOfType("string")).Return(&domain.User{}, errors.New("unknown error"))

		uService := services.NewUserService(r)

		vUser := &domain.User{
			Username: mockUser.Username,
			Password: "wrong-password",
		}

		err := uService.Validate(vUser)

		assert.Error(t, err)
	})
}
