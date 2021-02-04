package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	"github.com/jake-hansen/agora/api/handlers"
	"github.com/jake-hansen/agora/api/middleware"
	"github.com/jake-hansen/agora/services"
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

	// Create auth handler
	authService, err := services.BuildSimpleAuthService()
	if err != nil {
		panic(err)
	}
	handlers.NewAuthHandler(v1, authService)

	return router
}

func setGinEnvironment(env string) {
	if env == "prod" {
		gin.SetMode("release")
	}
}