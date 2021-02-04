package userservicemock

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

func Provide() *UserService {
	return &UserService{mock.Mock{}}
}

var (
	ProviderSet = wire.NewSet(Provide, wire.Bind(new(domain.UserService), new(*UserService)))
)