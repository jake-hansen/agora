package userservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// UserService is a mock of a UserService.
type UserService struct {
	mock.Mock
}

// Register mocks the Register function of UserService.
func (u *UserService) Register(user *domain.User) (uint, error) {
	args := u.Called(user)
	return uint(args.Int(0)), args.Error(1)
}

// Validate mocks the Validate function of UserService.
func (u *UserService) Validate(credentials *domain.Credentials) (*domain.User, error) {
	args := u.Called(credentials)
	return args.Get(0).(*domain.User), args.Error(1)
}

// GetAll mocks the GetAll function of UserService.
func (u *UserService) GetAll() ([]*domain.User, error) {
	args := u.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

// GetByID mocks the GetByID function of UserService.
func (u *UserService) GetByID(ID uint) (*domain.User, error) {
	args := u.Called(ID)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Update mocks the Update function of UserService.
func (u *UserService) Update(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}

// Delete mocks the Delete function of UserService.
func (u *UserService) Delete(ID uint) error {
	args := u.Called(ID)
	return args.Error(0)
}

func (u *UserService) GetByUsername(username string) (*domain.User, error) {
	args := u.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}
