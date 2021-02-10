package jwtservicemock

import (
	"github.com/google/wire"
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock JWT Service.
func Provide() *Service {
	return &Service{mock.Mock{}}
}

var (
	// ProviderSet provides a mock JWT Service for use in testing.
	ProviderSet = wire.NewSet(Provide)
)
