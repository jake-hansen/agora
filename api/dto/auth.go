package dto

// Credentials represents a username and password combination.
type Credentials struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// Auth contains the credentials needed to begin the authentication process.
type Auth struct {
	Credentials *Credentials `json:"credentials,omitempty" binding:"required"`
}

// Token contains a string based token that can be used for authentication.
type Token struct {
	Value string `json:"token,omitempty" binding:"required"`
}
