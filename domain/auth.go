package domain

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
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
	IsAuthenticated(token TokenValue) (bool, error)
	Authenticate(auth Auth) (*TokenSet, error)
	RefreshToken(token TokenValue) (*TokenSet, error)
	Deauthenticate(token TokenValue) error
	GetUserFromAuthToken(token TokenValue) (*User, error)
}

type AuthClaims struct {
	jwt.StandardClaims
	UserID	uint	`json:"user_id"`
	Usage 	string	`json:"usage"`
}

type AuthToken struct {
	Value TokenValue
	Expires time.Time
	JWTClaims	AuthClaims
}

type TokenValue string

type RefreshToken struct {
	gorm.Model
	Value           TokenValue `gorm:"column:token_hash"`
	ExpiresAt       time.Time
	TokenNonceHash  string
	ParentTokenHash string
	UserID          uint
	Revoked         bool
	JWTClaims       RefreshClaims	`gorm:"-"`
}

type RefreshClaims struct {
	jwt.StandardClaims
	UserID uint	`json:"user_id"`
	AuthTokenHash string	`json:"auth_token_hash"`
	ParentTokenHash string	`json:"parent_token_hash"`
	Nonce	string	`json:"nonce"`
	Usage 	string	`json:"usage"`
}

func (r RefreshToken) Hash() string {
	return r.Value.hash()
}

func (r TokenValue) hash() string {
	return hash(string(r))
}

func hash(v string) string {
	hasher := sha256.New()
	hasher.Write([]byte(v))
	value := hex.EncodeToString(hasher.Sum(nil))
	return value
}

func (r TokenValue) Value() (driver.Value, error) {
	return r.hash(), nil
}

func (r *TokenValue) Scan(src interface{}) error {
	bytes := src.([]byte)
	*r = TokenValue(bytes)
	return nil
}

type RefreshTokenRepository interface {
	Create(token RefreshToken) (uint, error)
	GetAll() ([]*RefreshToken, error)
	GetByToken(token RefreshToken) (*RefreshToken, error)
	GetByTokenHash(hash string) (*RefreshToken, error)
	GetByParentTokenHash(hash string, nonceHash string) (*RefreshToken, error)
	GetByTokenNonceHash(nonceHash string) (*RefreshToken, error)
	Update(token *RefreshToken) error
	Delete(ID uint) error
}

type RefreshTokenService interface {
	SaveNewRefreshToken(token RefreshToken) (uint, error)
	ReplaceRefreshToken(token RefreshToken) error
	GetRefreshTokenByParentTokenHash(token RefreshToken) (*RefreshToken, error)
	RevokeLatestRefreshTokenByNonce(token RefreshToken) error
}

type TokenSet struct {
	Auth    AuthToken
	Refresh RefreshToken
}
