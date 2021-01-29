package domain

// User represents information about a user.
type User struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

// NewUser returns a pointer to a User which contains the given parameters.
func NewUser(firstname string, lastname string, username string, password string) *User {
	user := &User{
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Password:  password,
	}
	return user
}
