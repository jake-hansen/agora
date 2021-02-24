package authservicemock

import (
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock AuthService.
func Provide() *AuthService {
	return &AuthService{mock.Mock{}}
}
