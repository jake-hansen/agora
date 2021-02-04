// +build wireinject

package database

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func BuildDB() *gorm.DB {
	panic(wire.Build(DBSet))
}
