package repositories

import (
	"fmt"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) Create(user *domain.User) (uint, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return user.ID, nil
}

func (u *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all users: %w", err)
	}
	return users, nil
}

func (u *UserRepository) GetByUsername(username string) (*domain.User, error) {
	// Need to put constraint on username to ensure it is unique
	user := new(domain.User)
	if err := u.DB.Where("username = ?", username).First(user).Error; err != nil {
		return nil, fmt.Errorf("error retrieving user by username %s: %w", username, err)
	}
	return user, nil
}

func (u *UserRepository) GetByID(ID uint) (*domain.User, error) {
	user := new(domain.User)
	if err := u.DB.First(user, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving user with id %d: %w", ID, err)
	}
	return user, nil
}

func (u *UserRepository) Update(user *domain.User) error {
	if err := u.DB.Model(user).Updates(domain.User{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Username:  user.Username,
		Password:  user.Password,
	}).Error; err != nil {
		return fmt.Errorf("error updating user with id %d: %w", user.ID, err)
	}
	return nil
}

func (u *UserRepository) Delete(ID uint) error {
	if err := u.DB.Delete(&domain.User{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting using with id %d: %w", ID, err)
	}
	return nil
}

var (
	UserRepositorySet = wire.NewSet(ProvideUserRepository, wire.Bind(new(domain.UserRepository), new(*UserRepository)), database.ProvideDB)
)
