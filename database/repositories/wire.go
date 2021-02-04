// +build wireinject

package repositories

import (
	"github.com/google/wire"
	"github.com/jake-hansen/agora/database"
)

func BuildUserRepository() *UserRepository {
	panic(wire.Build(UserRepositorySet, database.DBSet))
}