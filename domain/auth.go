package domain

type Auth struct {
	Credentials *Credentials
}

type Credentials struct {
	Username	string
	Password	string
}


type AuthService interface {
	IsAuthenticated(token Token) (bool, error)
	Authenticate(auth Auth) (*Token, error)
	Deauthenticate(token Token) error
}
