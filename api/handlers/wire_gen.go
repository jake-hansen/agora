// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package handlers

import (
	"github.com/jake-hansen/agora/router/handlers"
)

// Injectors from injector.go:

func Build() []handlers.Handler {
	v := ProvideAllProductionHandlers()
	return v
}
