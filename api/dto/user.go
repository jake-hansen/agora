package dto

// User represents information about a user.
type User struct {
	ID        uint   `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty" binding:"required"`
	Lastname  string `json:"lastname,omitempty" binding:"required"`
	Username  string `json:"username,omitempty" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
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
