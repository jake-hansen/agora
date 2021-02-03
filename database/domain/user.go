package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname	string
	Lastname	string
	Username	string
	Password	string
}

type UserRepository interface {
	Create(user *User) (uint, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}

type UserService interface {
	Create(user *User) (uint, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}
