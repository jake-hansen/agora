package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/domain"
	"time"
)

type Config struct {
	Issuer     string
	SigningKey string
	Duration   time.Duration
}

type Service struct {
	config Config
}

type claims struct {
	jwt.StandardClaims
}

// GenerateToken creates a JWT for the specified userrepo and returns the token as a string.
func (j *Service) GenerateToken(user domain.User) (string, error) {
	now := time.Now()

	claims := &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(j.config.Duration).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.config.Issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.config.SigningKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// ValidateToken validates the given token string. If the token is valid, the token string is return as a jwt.Token.
// Otherwise, a nil token is returned along with an error.
func (j *Service) ValidateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.config.SigningKey), nil
	})

	return t, err
}
