package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/api"
	handlers2 "github.com/jake-hansen/agora/router/handlers"
)

type Config struct {
	Environment		string
	Middleware		[]gin.HandlerFunc
	Handlers		[]handlers2.Handler
	RootEndpoint	string
}

type Router struct {
	engine		*gin.Engine
	config		*Config
}

func (r *Router) init() {
	r.engine = gin.New()
	r.setGinEnvironment(r.config.Environment)
	api.RegisterCustomValidation()

	for _, middleware := range r.config.Middleware {
		r.engine.Use(middleware)
	}

	parentGroup := r.engine.Group(r.config.RootEndpoint)
	versionGroup := parentGroup.Group("v1")

	for _, handler := range r.config.Handlers {
		err := handler.Register(versionGroup)
		if err != nil {
			fmt.Printf("could not register handler: %s\n", err.Error())
		}
	}
}

func (r *Router) Run(address string) error {
	err := r.engine.Run(address)
	return err
}

func New(cfg Config) *Router {
	r := &Router{config: &cfg}
	r.init()
	return r
}

func (r *Router) setGinEnvironment(env string) {
	if env == "prod" {
		gin.SetMode("release")
	}
}
