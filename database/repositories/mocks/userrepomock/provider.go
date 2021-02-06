package userrepomock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

func Provide() *UserRepository {
	return &UserRepository{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.UserRepository), new(*UserRepository)))
)
