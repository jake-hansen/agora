package handlers

import "github.com/gin-gonic/gin"

// Handler manages a subset of endpoints for a RouterGroup.
type Handler interface {
	Register(parentGroup *gin.RouterGroup) error
}

// HandlerManager manages a set of Handlers.
type HandlerManager struct {
	Handlers *[]Handler
}

func NewHandlerManager(handlers *[]Handler) *HandlerManager {
	return &HandlerManager{Handlers: handlers}
}
