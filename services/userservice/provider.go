package userservice

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
)

// Provide provides a new UserService configured with the given UserRepository.
func Provide(repository domain.UserRepository) *UserService {
	return &UserService{repo: repository}
}

var (
	// ProviderProductionSet provides a new UserService for use in production.
	ProviderProductionSet = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))

	// ProviderTestSet provides a new UserService for use in testing.
	ProviderTestSet = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))
)
