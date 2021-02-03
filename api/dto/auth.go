package dto

type Auth struct {
	Credentials *User  `json:"credentials,omitempty" binding:"required"`
}

type Token struct {
	Value	string `json:"token,omitempty" binding:"required"`
}
