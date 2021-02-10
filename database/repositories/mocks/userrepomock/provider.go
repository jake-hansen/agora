package userrepomock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

// Provide provides a UserRepository mock.
func Provide() *UserRepository {
	return &UserRepository{mock.Mock{}}
}

var (
	//ProviderSet provides an instance of a UserRepository mock.
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.UserRepository), new(*UserRepository)))
)
