package domain

type Auth struct {
	Credentials *User  `json:"credentials,omitempty"`
	Token       string `json:"token,omitempty"`
}

type AuthService interface {
	IsAuthenticated(auth Auth) (bool, error)
	Authenticate(auth Auth) (interface{}, error)
	Deauthenticate(auth Auth) error
}
