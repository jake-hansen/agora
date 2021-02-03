package repository_mocks

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) Create(user *domain.User) (uint, error) {
	args := u.Called(user)
	return uint(args.Int(0)), args.Error(1)
}

func (u *UserRepository) GetAll() ([]*domain.User, error) {
	args := u.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (u *UserRepository) GetByID(ID uint) (*domain.User, error) {
	args := u.Called(ID)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepository) GetByUsername(username string) (*domain.User, error) {
	args := u.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserRepository) Update(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserRepository) Delete(ID uint) error {
	args := u.Called(ID)
	return args.Error(0)
}

