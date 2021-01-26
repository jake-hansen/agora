package domain

// User represents information about a user.
type User struct {
	Firstname	string
	Lastname	string
	Username	string
	Password	string
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
