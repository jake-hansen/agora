package userrepo

import (
	"fmt"
	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (u *UserRepository) Create(user *domain.User) (uint, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("error creating userrepo: %w", err)
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
		return nil, fmt.Errorf("error retrieving userrepo by username %s: %w", username, err)
	}
	return user, nil
}

func (u *UserRepository) GetByID(ID uint) (*domain.User, error) {
	user := new(domain.User)
	if err := u.DB.First(user, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving userrepo with id %d: %w", ID, err)
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
		return fmt.Errorf("error updating userrepo with id %d: %w", user.ID, err)
	}
	return nil
}

func (u *UserRepository) Delete(ID uint) error {
	if err := u.DB.Delete(&domain.User{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting using with id %d: %w", ID, err)
	}
	return nil
}
