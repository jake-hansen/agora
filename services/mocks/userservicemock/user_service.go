package userservicemock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

type UserService struct {
	mock.Mock
}

func (u *UserService) Register(user *domain.User) (uint, error) {
	args := u.Called(user)
	return uint(args.Int(0)), args.Error(1)
}

func (u *UserService) Validate(credentials *domain.Credentials) (*domain.User, error) {
	args := u.Called(credentials)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserService) GetAll() ([]*domain.User, error) {
	args := u.Called()
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (u *UserService) GetByID(ID uint) (*domain.User, error) {
	args := u.Called(ID)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (u *UserService) Update(user *domain.User) error {
	args := u.Called(user)
	return args.Error(0)
}

func (u *UserService) Delete(ID uint) error {
	args := u.Called(ID)
	return args.Error(0)
}
