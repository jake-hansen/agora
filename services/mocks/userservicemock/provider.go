package userservicemock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// Provide provides a mock UserService.
func Provide() *UserService {
	return &UserService{mock.Mock{}}
}

var (
	// ProviderSet provides a mock UserService for use in testing.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))
)
