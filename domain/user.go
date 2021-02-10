package domain

import (
	"database/sql/driver"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname 	  string
	Lastname  	  string
	Username  	  string
	Password  	  *Password
}

type Password struct {
	plaintext	string
	Hash		[]byte
}

func NewPassword(plaintext string) *Password {
	return &Password{plaintext: plaintext}
}

func (p *Password) HashPassword() ([]byte, error) {
	pHash, err := bcrypt.GenerateFromPassword([]byte(p.plaintext), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return pHash, err
}

func (p *Password) Scan(src interface{}) error {
	p.Hash = make([]byte, len(src.([]byte)))
	copy(p.Hash, src.([]byte))
	return nil
}

func (p *Password) Value() (driver.Value, error) {
	return p.Hash, nil
}

type UserRepository interface {
	Create(user *User) (uint, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}

type UserService interface {
	Register(user *User) (uint, error)
	Validate(credentials *Credentials) (*User, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}
