package userservicemock

import (
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock UserService.
func Provide() *UserService {
	return &UserService{mock.Mock{}}
}
