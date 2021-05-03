package userrepo

import (
	"fmt"

	"github.com/jake-hansen/agora/database/repositories"

	"github.com/jake-hansen/agora/domain"
	"gorm.io/gorm"
)

// UserRepository is a repository that holds information about Users
// backed by a database.
type UserRepository struct {
	DB *gorm.DB
}

// Create creates the given user in the database.
func (u *UserRepository) Create(user *domain.User) (uint, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}
	return user.ID, nil
}

// GetAll gets all of the users in the database.
func (u *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all users: %w", err)
	}
	return users, nil
}

// GetByUsername retrieves a User based on the given username.
func (u *UserRepository) GetByUsername(username string) (*domain.User, error) {
	// Need to put constraint on username to ensure it is unique
	user := new(domain.User)
	if err := u.DB.Where("username = ?", username).First(user).Error; err != nil {
		return nil, repositories.NewNotFoundError(repositories.DATABASE_ACTION_RETRIEVE, "user", username, "by username")
	}
	return user, nil
}

// GetByID retrieves a User based on the given ID.
func (u *UserRepository) GetByID(ID uint) (*domain.User, error) {
	user := new(domain.User)
	if err := u.DB.First(user, ID).Error; err != nil {
		return nil, fmt.Errorf("error retrieving user with id %d: %w", ID, err)
	}
	return user, nil
}

// Update updates the given User. The ID of the given User needs to be set
// in order to find the existing record in the database.
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

// Delete deletes the User with the given ID.
func (u *UserRepository) Delete(ID uint) error {
	if err := u.DB.Delete(&domain.User{}, ID).Error; err != nil {
		return fmt.Errorf("error deleting user with id %d: %w", ID, err)
	}
	return nil
}
