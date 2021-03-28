package domain

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
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

type RefreshTokenValue string

type RefreshToken struct {
	gorm.Model
	Value           RefreshTokenValue	`gorm:"column:token_hash"`
	ExpiresAt       time.Time
	TokenNonceHash  string
	ParentTokenHash string
	UserID			uint
	Revoked			bool
}

func (r RefreshToken) Hash() string {
	return r.Value.hash()
}

func (r RefreshTokenValue) hash() string {
	hasher := sha256.New()
	hasher.Write([]byte(r))
	value := hex.EncodeToString(hasher.Sum(nil))
	return value
}

func (r RefreshTokenValue) Value() (driver.Value, error) {
	return r.hash(), nil
}

func (r *RefreshTokenValue) Scan(src interface{}) error {
	bytes := src.([]byte)
	*r = RefreshTokenValue(bytes)
	return nil
}

type RefreshTokenRepository interface {
	Create(token RefreshToken) (uint, error)
	GetAll() ([]*RefreshToken, error)
	GetByToken(token RefreshToken) (*RefreshToken, error)
	GetByTokenHash(hash string) (*RefreshToken, error)
	GetByParentTokenHash(hash string) (*RefreshToken, error)
	GetByTokenNonceHash(nonceHash string) (*RefreshToken, error)
	Update(token *RefreshToken) error
	Delete(ID uint) error
}

type RefreshTokenService interface {
	SaveNewRefreshToken(token RefreshToken) (uint, error)
	GetRefreshTokenByHash(hash string) (*RefreshToken, error)
	ReplaceRefreshToken(token RefreshToken) error
	GetRefreshTokenByParentTokenHash(hash string) (*RefreshToken, error)
	RevokeRefreshTokenByNonce(nonceHash string) error
}

type TokenSet struct {
	Auth    AuthToken
	Refresh RefreshToken
}
