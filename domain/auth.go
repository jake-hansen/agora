package domain

import (
	"gorm.io/gorm"
	"time"
)

// Auth contains the credentials needed to begin the authentication process.
type Auth struct {
	Credentials *Credentials
}

// Credentials represents a username and password combination.
type Credentials struct {
	Username string
	Password string
}

// AuthService manages authentication based on Auths and Tokens.
type AuthService interface {
	IsAuthenticated(token AuthToken) (bool, error)
	Authenticate(auth Auth) (*TokenSet, error)
	RefreshToken(token RefreshToken) (*TokenSet, error)
	Deauthenticate(token AuthToken) error
	GetUserFromAuthToken(token AuthToken) (*User, error)
}

type AuthToken struct {
	Value string
	Expires time.Time
}

type RefreshToken struct {
	gorm.Model
	Value string
	Expires time.Time
}

type TokenSet struct {
	Auth    AuthToken
	Refresh RefreshToken
}
