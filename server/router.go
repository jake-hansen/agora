package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/api/services"
	"time"
)

// NewRouter returns a router configured with handlers for configured
// endpoints.
func NewRouter(env string) *gin.Engine {
	setGinEnvironment(env)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.PublicErrorHandler())
	api.RegisterCustomValidation()

	v1 := router.Group("v1")
	dur, _ := time.ParseDuration("15m")
	handlers.NewAuthHandler(v1, services.NewSimpleAuthService(services.NewJWTService("agora", "test", dur)))

	return router
}

func setGinEnvironment(env string) {
	if env == "prod" {
		gin.SetMode("release")
	}
}