package authservicemock

import (
	"github.com/google/wire"
	"github.com/stretchr/testify/mock"
)

func Provide() *AuthService  {
	return &AuthService{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide)
)