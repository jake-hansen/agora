package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func ProvideAllProductionMiddleware() []gin.HandlerFunc {
	var middlewares []gin.HandlerFunc

	middlewares = append(middlewares, gin.Logger())
	middlewares = append(middlewares, gin.Recovery())
	middlewares = append(middlewares, PublicErrorHandler())

	return middlewares
}

var (
	ProviderProductionSet = wire.NewSet(ProvideAllProductionMiddleware)
)
