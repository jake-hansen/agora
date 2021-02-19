package middleware

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jake-hansen/agora/api/middleware/corsmiddleware"
	"github.com/jake-hansen/agora/log"
	"time"
)

// ProvideAllProductionMiddleware provides all the middleware that will be used in production.
func ProvideAllProductionMiddleware(log *log.Log, cors *corsmiddleware.CORSMiddleware) []gin.HandlerFunc {
	var middlewares []gin.HandlerFunc

	g := ginzap.Ginzap(log.ZapLogger, time.RFC3339, true)

	middlewares = append(middlewares, g)
	middlewares = append(middlewares, gin.Recovery())
	middlewares = append(middlewares, PublicErrorHandler())
	middlewares = append(middlewares, cors.HandleCORS())

	return middlewares
}

var (
	// ProviderProductionSet provides all middleware for production.
	ProviderProductionSet = wire.NewSet(ProvideAllProductionMiddleware, corsmiddleware.ProviderProductionSet)
)
