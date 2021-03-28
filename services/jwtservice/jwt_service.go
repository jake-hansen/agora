package jwtservice

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
)

// JWTService is a service for generating and validating JWTs.
type JWTService interface {
	GenerateAuthToken(user domain.User) (*domain.AuthToken, error)
	GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshTokenValue, expiry *time.Time) (*domain.RefreshToken, error)
	ValidateAuthToken(token string) (*jwt.Token, *AuthClaims, error)
	ValidateRefreshToken(token domain.RefreshTokenValue) (*jwt.Token, *RefreshClaims, error)
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

type AuthClaims struct {
	jwt.StandardClaims
	UserID	uint	`json:"user_id"`
}

type RefreshClaims struct {
	jwt.StandardClaims
	UserID uint	`json:"user_id"`
	AuthTokenHash string	`json:"auth_token_hash"`
	ParentTokenHash string	`json:"parent_token_hash"`
	Nonce	string	`json:"nonce"`
}

// GenerateToken creates a JWT for the specified User and returns the token as a string.
func (j *JWTServiceImpl) GenerateAuthToken(user domain.User) (*domain.AuthToken, error) {
	now := time.Now()
	expiry := now.Add(j.config.Duration)

	claims := &AuthClaims{
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

	domainToken := &domain.AuthToken{
		Value:   t,
		Expires: expiry,
	}

	return domainToken, nil
}

func (j *JWTServiceImpl) GenerateRefreshToken(user domain.User, authToken domain.AuthToken, parentToken *domain.RefreshTokenValue, expiry *time.Time) (*domain.RefreshToken, error) {
	now := time.Now()

	newExpiry := now.Add(j.config.RefreshDuration)
	if expiry != nil {
		newExpiry = *expiry
	}

	hasher := sha256.New()
	hasher.Write([]byte(authToken.Value))
	authSHA := hex.EncodeToString(hasher.Sum(nil))

	parentRefreshSHA := ""
	nonce := ""
	if parentToken != nil {
		hasher.Reset()
		hasher.Write([]byte(*parentToken))
		parentRefreshSHA = hex.EncodeToString(hasher.Sum(nil))

		parentNonce, err := j.getNonce(string(*parentToken))
		if err != nil {
			return nil, err
		}
		nonce = parentNonce
	} else {
		generatedNonce, err := j.generateNonce()
		if err != nil {
			return nil, err
		}
		nonce = generatedNonce
	}

	hasher.Reset()
	hasher.Write([]byte(nonce))
	nonceSHA := hex.EncodeToString(hasher.Sum(nil))

	claims := &RefreshClaims{
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
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return nil, err
	}

	domainToken := &domain.RefreshToken{
		Value:           domain.RefreshTokenValue(t),
		ExpiresAt:       newExpiry,
		ParentTokenHash: parentRefreshSHA,
		TokenNonceHash:  nonceSHA,
		UserID: user.ID,
	}

	return domainToken, nil
}

// ValidateToken validates the given token string. If the token is valid, the token string is return as a jwt.Token.
// Otherwise, a nil token is returned along with an error.
func (j *JWTServiceImpl) ValidateAuthToken(token string) (*jwt.Token, *AuthClaims, error) {
	var returnClaims *AuthClaims
	t, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tclaims, ok := token.Claims.(*AuthClaims); ok {
			returnClaims = tclaims
		}

		return []byte(j.config.SigningKey), nil
	})

	return t, returnClaims, err
}

func (j *JWTServiceImpl) ValidateRefreshToken(token domain.RefreshTokenValue) (*jwt.Token, *RefreshClaims, error) {
	var returnClaims *RefreshClaims
	t, err := jwt.ParseWithClaims(string(token), &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		if tclaims, ok := token.Claims.(*RefreshClaims); ok {
			returnClaims = tclaims
		}

		return []byte(j.config.SigningKey), nil
	})

	return t, returnClaims, err
}

func (j *JWTServiceImpl) getNonce(token string) (string, error)  {
	_, claims, err := j.ValidateRefreshToken(domain.RefreshTokenValue(token))

	if claims != nil {
		return claims.Nonce, nil
	} else {
		return "", err
	}
}

func (j *JWTServiceImpl) generateNonce() (string, error) {
	b := make([]byte, 64)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
