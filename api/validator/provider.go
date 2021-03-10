package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/google/wire"
)

func Cfg(customValidationFuncs []CustomValidationFunc) Config {
	c := Config {
		Engine:                binding.Validator.Engine(),
		CustomValidationFuncs: customValidationFuncs,
	}
	return c
}

func Provide(config Config) (*Validator, error) {
	return NewValidator(config)
}

func ProvideCustomValidationFuncs() []CustomValidationFunc {
	var funcs []CustomValidationFunc
	funcs = append(funcs, MeetingTimeValidator)
	return funcs
}

var (
	ProviderSet = wire.NewSet(Provide, Cfg, ProvideCustomValidationFuncs)
)
