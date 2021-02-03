package dto

type Auth struct {
	Credentials *User  `json:"credentials,omitempty" binding:"required"`
}

type Token struct {
	Value	string `json:"token,omitempty" binding:"required"`
}

type AuthService interface {
	IsAuthenticated(token Token) (bool, error)
	Authenticate(auth Auth) (*Token, error)
	Deauthenticate(token Token) error
}
