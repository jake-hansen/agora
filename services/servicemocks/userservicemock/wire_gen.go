// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package userservicemock

// Injectors from injector.go:

func Build() *UserService {
	userService := Provide()
	return userService
}
