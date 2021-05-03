package domain

import (
	"database/sql/driver"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user.
type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Username  string
	Password  *Password
}

// Password represents both a plaintext and hased password.
type Password struct {
	plaintext string
	Hash      []byte
}

// NewPassword creates a new Password based on the plaintext given.
func NewPassword(plaintext string) *Password {
	return &Password{plaintext: plaintext}
}

// HashPassword hashes and salts the password.
func (p *Password) HashPassword() ([]byte, error) {
	pHash, err := bcrypt.GenerateFromPassword([]byte(p.plaintext), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return pHash, err
}

// Scan takes the given src as a []byte and stores it in Hash.
func (p *Password) Scan(src interface{}) error {
	p.Hash = make([]byte, len(src.([]byte)))
	copy(p.Hash, src.([]byte))
	return nil
}

// Value returns the stored Hash of the password.
func (p *Password) Value() (driver.Value, error) {
	return p.Hash, nil
}

// UserRepository manages storing and retrieving information about a User.
type UserRepository interface {
	Create(user *User) (uint, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}

// UserService manages processing information about a User.
type UserService interface {
	Register(user *User) (uint, error)
	Validate(credentials *Credentials) (*User, error)
	GetAll() ([]*User, error)
	GetByID(ID uint) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(ID uint) error
}
