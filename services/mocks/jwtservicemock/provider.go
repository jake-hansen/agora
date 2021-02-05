package jwtservicemock

import (
	"github.com/google/wire"
	"github.com/stretchr/testify/mock"
)

func Provide() *Service {
	return &Service{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide)
)
