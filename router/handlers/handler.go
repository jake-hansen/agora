package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(parentGroup *gin.RouterGroup) error
}

type HandlerManager struct {
	Handlers	[]Handler
}

