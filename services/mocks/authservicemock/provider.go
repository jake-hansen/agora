package authservicemock

import (
	"github.com/google/wire"
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock AuthService.
func Provide() *AuthService {
	return &AuthService{mock.Mock{}}
}

var (
	// ProviderSet provides a mock AuthService for use in testing.
	ProviderSet = wire.NewSet(Provide)
)
