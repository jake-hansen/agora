package domain

type Auth struct {
	User	User	`json:"user,omitempty"`
	Token	string	`json:"token,omitempty"`
}

type AuthService interface {
	IsAuthenticated(auth Auth) (bool, error)
	Authenticate(auth Auth) (interface{}, error)
	Deauthenticate(auth Auth) error
}
