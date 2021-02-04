package services

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func ProvideUserService(repository domain.UserRepository) *UserService {
	return &UserService{repo: repository}
}

// Register creates a new User in the database. Note that the given User's password
// is first hashed before being saved in the repository.
func (u *UserService) Register(user *domain.User) (uint, error) {
	pHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error hashing password during registration: %w", err)
	}

	newUser := *user
	newUser.Password = string(pHash)

	return u.repo.Create(&newUser)
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

func (u *UserService) Validate(user *domain.User) error {
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

var (
	UserServiceSet = wire.NewSet(ProvideUserService, wire.Bind(new(domain.UserService), new(*UserService)))
)

