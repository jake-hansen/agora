package domain

// Auth contains the credentials needed to begin the authentication process.
type Auth struct {
	Credentials *Credentials
}

// Credentials represents a username and password combination.
type Credentials struct {
	Username string
	Password string
}

// AuthService manages authentication based on Auths and Tokens.
type AuthService interface {
	IsAuthenticated(token Token) (bool, error)
	Authenticate(auth Auth) (*Token, error)
	Deauthenticate(token Token) error
}

// Token contains a string based token that can be used for authentication.
type Token struct {
	Value string
}
