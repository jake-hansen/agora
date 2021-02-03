package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jake-hansen/agora/api/dto"
	"time"
)

type JWTService struct {
	issuer		string
	signingKey	string
	jwtDuration	time.Duration
}

type claims struct {
	jwt.StandardClaims
}

// NewJWTService returns a new JWTService with the specified parameters.
func NewJWTService(issuer string, signingKey string, jwtDuration time.Duration) JWTService {
	return JWTService{
		issuer:     	issuer,
		signingKey: 	signingKey,
		jwtDuration: 	jwtDuration,
	}
}

// GenerateToken creates a JWT for the specified user and returns the token as a string.
func (j *JWTService) GenerateToken(user dto.User) (string, error) {
	now := time.Now()

	claims := &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(j.jwtDuration).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    j.issuer,
			NotBefore: now.Unix(),
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

// ValidateToken validates the given token string. If the token is valid, the token string is return as a jwt.Token.
// Otherwise, a nil token is returned along with an error.
func (j *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.signingKey), nil
	})

	return t, err
}
