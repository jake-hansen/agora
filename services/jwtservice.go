package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/spf13/viper"
	"time"
)

type JWTService struct {
	config JWTConfig
}

type JWTConfig struct {
	Issuer     string
	SigningKey string
	Duration   time.Duration
}

func ProvideJWTConfig(v *viper.Viper) (*JWTConfig, error) {
	dur, err := time.ParseDuration(v.GetString("jwtservice.duration"))
	if err != nil {
		return nil, err
	}

	cfg := &JWTConfig{
		Issuer:     v.GetString("jwtservice.issuer"),
		SigningKey: v.GetString("jwtservice.signingkey"),
		Duration:   dur,
	}

	return cfg, nil
}

type claims struct {
	jwt.StandardClaims
}

// ProvideJWTService returns a new JWTService with the specified config.
func ProvideJWTService(config *JWTConfig) *JWTService {
	return &JWTService{*config}
}

// GenerateToken creates a JWT for the specified user and returns the token as a string.
func (j *JWTService) GenerateToken(user domain.User) (string, error) {
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
func (j *JWTService) ValidateToken(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.config.SigningKey), nil
	})

	return t, err
}

var (
	JWTServiceSet = wire.NewSet(ProvideJWTService, ProvideJWTConfig)
)
