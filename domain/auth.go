package domain

import "github.com/jake-hansen/agora/api/dto"

type Auth struct {
	Credentials *User
}

type Token struct {
	Value	string
}

type AuthService interface {
	IsAuthenticated(token dto.Token) (bool, error)
	Authenticate(auth dto.Auth) (*dto.Token, error)
	Deauthenticate(token dto.Token) error
}
