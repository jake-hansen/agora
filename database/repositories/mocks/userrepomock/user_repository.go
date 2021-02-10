package userrepomock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// UserRepository is a mock of a UserRepository.
type UserRepository struct {
	mock.Mock
}

// Create mocks the Create function.
func (u *UserRepository) Create(user *domain.User) (uint, error) {
	args := u.Called(user)
	return uint(args.Int(0)), args.Error(1)
}

// GetAll mocks the GetAll function.
func (u *UserRepository) GetAll() ([]*domain.User, error) {
	args := u.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

// GetByID mocks the GetByID function.
func (u *UserRepository) GetByID(ID uint) (*domain.User, error) {
	args := u.Called(ID)
	return args.Get(0).(*domain.User), args.Error(1)
}

// GetByUsername mocks the GetByUsername function.
func (u *UserRepository) GetByUsername(username string) (*domain.User, error) {
	args := u.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Update mocks the Update function.
func (u *UserRepository) Update(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}

// Delete mocks the Delete function.
func (u *UserRepository) Delete(ID uint) error {
	args := u.Called(ID)
	return args.Error(0)
}
