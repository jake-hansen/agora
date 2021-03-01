package userrepomock

import (
	"github.com/stretchr/testify/mock"
)

// Provide provides a UserRepository mock.
func Provide() *UserRepository {
	return &UserRepository{mock.Mock{}}
}
