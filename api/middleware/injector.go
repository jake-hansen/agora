// +build wireinject

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func Build() []gin.HandlerFunc {
	panic(wire.Build(ProviderProductionSet))
}