// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jake-hansen/agora/log"
)

// Injectors from injector.go:

func Build(log2 *log.Log) []gin.HandlerFunc {
	v := ProvideAllProductionMiddleware(log2)
	return v
}
