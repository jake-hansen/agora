package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/google/wire"
)

// Cfg returns a new Config containing the given slice of CustomValidationFuncs.
func Cfg(customValidationFuncs []CustomValidationFunc) Config {
	c := Config{
		Engine:                binding.Validator.Engine(),
		CustomValidationFuncs: customValidationFuncs,
	}
	return c
}

// Provide provides a new Validator using the given Config.
func Provide(config Config) (*Validator, error) {
	return NewValidator(config)
}

// ProvideCustomValidationFuncs provides a slice of CustomValidationFuncs used in the application.
func ProvideCustomValidationFuncs() []CustomValidationFunc {
	var funcs []CustomValidationFunc
	funcs = append(funcs, MeetingTimeValidator)
	return funcs
}

var (
	// ProviderSet provides a Validator.
	ProviderSet = wire.NewSet(Provide, Cfg, ProvideCustomValidationFuncs)
)
