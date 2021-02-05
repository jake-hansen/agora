package dto

type Credentials struct {
	Username	string	`json:"username,omitempty" binding:"required"`
	Password	string	`json:"password,omitempty" binding:"required"`
}

type Auth struct {
	Credentials *Credentials  `json:"credentials,omitempty" binding:"required"`
}

type Token struct {
	Value	string `json:"token,omitempty" binding:"required"`
}
