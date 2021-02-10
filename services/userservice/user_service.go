package userservice

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

// Register creates a new User in the database. Note that the given User's password
// is first hashed before being saved in the repository.
func (u *UserService) Register(user *domain.User) (uint, error) {
	return u.repo.Create(user)
}

func (u *UserService) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}

func (u *UserService) GetByID(ID uint) (*domain.User, error) {
	return u.repo.GetByID(ID)
}

func (u *UserService) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *UserService) Delete(ID uint) error {
	return u.repo.Delete(ID)
}

func (u *UserService) Validate(credentials *domain.Credentials) (*domain.User, error) {
	errMsg := "could not validate user: %w"
	foundUser, err := u.repo.GetByUsername(credentials.Username)
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password.Hash), []byte(credentials.Password))
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}
	return foundUser, nil
}
