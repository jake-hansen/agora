package jwtservice

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
)

// JWTService is a service for generating and validating JWTs.
type JWTService interface {
	GenerateToken(user domain.User) (*domain.Token, error)
	GenerateRefreshToken(user domain.User, authToken domain.Token) (*domain.Token, error)
	ValidateToken(token string) (*jwt.Token, *Claims, error)
}

// Config contains the parameters for configuring a JWTService.
type Config struct {
	Issuer     string
	SigningKey string
	Duration   time.Duration
	RefreshDuration time.Duration
}

// JWTServiceImpl is an implementation of a JWTService.
type JWTServiceImpl struct {
	config Config
}

type Claims struct {
	jwt.StandardClaims
	UserID	uint
}

type RefreshClaims struct {
	jwt.StandardClaims
	UserID uint
	AuthTokenHash string
}

// GenerateToken creates a JWT for the specified User and returns the token as a string.
func (j *JWTServiceImpl) GenerateToken(user domain.User) (*domain.Token, error) {
	now := time.Now()
	expiry := now.Add(j.config.Duration)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.config.Issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return nil, err
	}

	domainToken := &domain.Token{
		Value:   t,
		Expires: expiry,
	}

	return domainToken, nil
}

func (j *JWTServiceImpl) GenerateRefreshToken(user domain.User, authToken domain.Token) (*domain.Token, error) {
	now := time.Now()
	expiry := now.Add(j.config.RefreshDuration)

	hasher := sha256.New()
	hasher.Write([]byte(authToken.Value))
	sha := hex.EncodeToString(hasher.Sum(nil))

	claims := &RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.config.Issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
		UserID:         user.ID,
		AuthTokenHash:  sha,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return nil, err
	}

	domainToken := &domain.Token{
		Value:   t,
		Expires: expiry,
	}

	return domainToken, nil
}

// ValidateToken validates the given token string. If the token is valid, the token string is return as a jwt.Token.
// Otherwise, a nil token is returned along with an error.
func (j *JWTServiceImpl) ValidateToken(token string) (*jwt.Token, *Claims, error) {
	var returnClaims *Claims
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tclaims, ok := token.Claims.(*Claims); ok {
			returnClaims = tclaims
		}

		return []byte(j.config.SigningKey), nil
	})

	return t, returnClaims, err
}
