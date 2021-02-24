package jwtservicemock

import (
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock JWT Service.
func Provide() *Service {
	return &Service{mock.Mock{}}
}
