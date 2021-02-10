package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// ProvideAllProductionMiddleware provides all the middleware that will be used in production.
func ProvideAllProductionMiddleware() []gin.HandlerFunc {
	var middlewares []gin.HandlerFunc

	middlewares = append(middlewares, gin.Logger())
	middlewares = append(middlewares, gin.Recovery())
	middlewares = append(middlewares, PublicErrorHandler())

	return middlewares
}

var (
	// ProviderProductionSet provides all middleware for production.
	ProviderProductionSet = wire.NewSet(ProvideAllProductionMiddleware)
)
