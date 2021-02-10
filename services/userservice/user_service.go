package userservice

import (
	"fmt"

	"github.com/jake-hansen/agora/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserService is a service which processes information about a User.
type UserService struct {
	repo domain.UserRepository
}

// Register creates a new User in the repository.
func (u *UserService) Register(user *domain.User) (uint, error) {
	return u.repo.Create(user)
}

// GetAll retrieves all Users in the repository.
func (u *UserService) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}

// GetByID retrieves the User with the given ID from the repository.
func (u *UserService) GetByID(ID uint) (*domain.User, error) {
	return u.repo.GetByID(ID)
}

// Update updates the given User in the repository.
func (u *UserService) Update(user *domain.User) error {
	return u.repo.Update(user)
}

// Delete deletes the given User in the repository.
func (u *UserService) Delete(ID uint) error {
	return u.repo.Delete(ID)
}

// Validate validates the given credentials by comparing the given plaintext password
// with the hashed password in the repository for the given username. If the credentials
// are valid, the User is returned. Otherwise, an error is returned with a nil User.
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
