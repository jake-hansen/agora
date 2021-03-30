package jwtservice

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
)

const (
	USAGE_AUTH = "auth"
	USAGE_REFRESH = "refresh"
)

// JWTService is a service for generating and validating JWTs.
type JWTService interface {
	GenerateAuthToken(user domain.User) (*domain.AuthToken, error)
	GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshToken) (*domain.RefreshToken, error)
	ValidateAuthToken(token domain.TokenValue) (domain.AuthToken, error)
	ValidateRefreshToken(token domain.TokenValue) (domain.RefreshToken, error)
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

// GenerateToken creates a JWT for the specified User and returns the token as a string.
func (j *JWTServiceImpl) GenerateAuthToken(user domain.User) (*domain.AuthToken, error) {
	now := time.Now()
	expiry := now.Add(j.config.Duration)

	claims := &domain.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.config.Issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
		UserID: user.ID,
		Usage: USAGE_AUTH,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return nil, err
	}

	domainToken := &domain.AuthToken{
		Value:   domain.TokenValue(t),
		Expires: expiry,
		JWTClaims: *claims,
	}

	return domainToken, nil
}

func (j *JWTServiceImpl) GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshToken) (*domain.RefreshToken, error) {
	now := time.Now()

	newExpiry := now.Add(j.config.RefreshDuration)
	if parentToken != nil {
		newExpiry = parentToken.ExpiresAt
	}

	authSHA := hash(string(authToken.Value))

	parentRefreshSHA := ""
	nonce := ""
	nonceSHA := ""
	if parentToken != nil {
		parentRefreshSHA = parentToken.Hash()
		nonce = parentToken.JWTClaims.Nonce
		nonceSHA = parentToken.TokenNonceHash
	} else {
		generatedNonce, err := j.generateNonce()
		if err != nil {
			return nil, err
		}
		nonce = generatedNonce
		nonceSHA = hash(nonce)
	}

	claims := &domain.RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: newExpiry.Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.config.Issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
		UserID:        user.ID,
		AuthTokenHash: authSHA,
		ParentTokenHash: parentRefreshSHA,
		Nonce: nonce,
		Usage: USAGE_REFRESH,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return nil, err
	}

	domainToken := &domain.RefreshToken{
		Value:           domain.TokenValue(t),
		ExpiresAt:       newExpiry,
		ParentTokenHash: parentRefreshSHA,
		TokenNonceHash:  nonceSHA,
		UserID: user.ID,
		JWTClaims: *claims,
	}

	return domainToken, nil
}

// ValidateToken validates the given token string. If the token is valid, the token string is return as a jwt.Token.
// Otherwise, a nil token is returned along with an error.
func (j *JWTServiceImpl) ValidateAuthToken(token domain.TokenValue) (domain.AuthToken, error) {
	var claims *domain.AuthClaims
	_, err := jwt.ParseWithClaims(string(token), &domain.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tclaims, ok := token.Claims.(*domain.AuthClaims); ok {
			if tclaims.Usage != USAGE_AUTH {
				return nil, errors.New("token is not an auth token")
			}
			claims = tclaims
		}

		return []byte(j.config.SigningKey), nil
	})
	if err != nil {
		return domain.AuthToken{}, err
	}

	returnToken := domain.AuthToken{
		Value:     token,
		Expires:   time.Unix(claims.ExpiresAt, 0),
		JWTClaims: *claims,
	}

	return returnToken, err
}

func (j *JWTServiceImpl) ValidateRefreshToken(token domain.TokenValue) (domain.RefreshToken, error) {
	var claims *domain.RefreshClaims
	_, err := jwt.ParseWithClaims(string(token), &domain.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tclaims, ok := token.Claims.(*domain.RefreshClaims); ok {
			if tclaims.Usage != USAGE_REFRESH {
				return nil, errors.New("token is not an refresh token")
			}
			claims = tclaims
		}

		return []byte(j.config.SigningKey), nil
	})
	if err != nil {
		return domain.RefreshToken{}, err
	}

	returnToken := domain.RefreshToken{
		Value:           token,
		ExpiresAt:       time.Unix(claims.ExpiresAt, 0),
		TokenNonceHash:  hash(claims.Nonce),
		ParentTokenHash: claims.ParentTokenHash,
		UserID:          claims.UserID,
		Revoked:         false,
		JWTClaims:       *claims,
	}

	return returnToken, err
}

func (j *JWTServiceImpl) generateNonce() (string, error) {
	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func hash(v string) string {
	hasher := sha256.New()
	hasher.Write([]byte(v))
	value := hex.EncodeToString(hasher.Sum(nil))
	return value
}
