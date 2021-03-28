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
	IsAuthenticated(token Token) (bool, error)
	Authenticate(auth Auth) (*TokenSet, error)
	RefreshToken(tokens TokenSet) (*TokenSet, error)
	Deauthenticate(token Token) error
	GetUser(token Token) (*User, error)
}

type Token struct {
	Value string
	Expires time.Time
}

type RefreshToken struct {
	gorm.Model
	Value string
	Expires time.Time
}

type TokenSet struct {
	Auth    Token
	Refresh Token
}
