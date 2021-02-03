package domain

type Auth struct {
	Credentials *User
}

type AuthService interface {
	IsAuthenticated(token Token) (bool, error)
	Authenticate(auth Auth) (*Token, error)
	Deauthenticate(token Token) error
}
