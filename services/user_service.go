package services

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository) domain.UserService {
	return &userService{repo: repository}
}

// Register creates a new User in the database. Note that the given User's password
// is first hashed before being saved in the repository.
func (u *userService) Register(user *domain.User) (uint, error) {
	pHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error hashing password during registration: %w", err)
	}

	newUser := *user
	newUser.Password = string(pHash)

	return u.repo.Create(&newUser)
}

func (u *userService) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}

func (u *userService) GetByID(ID uint) (*domain.User, error) {
	return u.repo.GetByID(ID)
}

func (u *userService) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userService) Delete(ID uint) error {
	return u.repo.Delete(ID)
}

func (u *userService) Validate(user *domain.User) error {
	errMsg := "could not validate user: %w"
	foundUser, err := u.repo.GetByUsername(user.Username)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	return nil
}

