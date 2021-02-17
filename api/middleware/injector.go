// +build wireinject

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/log"
)

func Build(log *log.Log) []gin.HandlerFunc {
	panic(wire.Build(ProviderProductionSet))
}