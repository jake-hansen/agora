package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type ValidationEngine interface{}

type CustomValidationFunc struct {
	Tag  string
	Func validator.Func
	CallValidationEvenIfNull bool
}

type Config struct {
	Engine                ValidationEngine
	CustomValidationFuncs []CustomValidationFunc
}

type Validator struct {
	engine 				  ValidationEngine
	customValidationFuncs []CustomValidationFunc
}

func NewValidator(config Config) (*Validator, error) {
	v := &Validator{
		engine: config.Engine,
	}
	err := v.init()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *Validator) init() error {
	err := v.registerCustomTagNameFunc()
	if err != nil {
		return err
	}
	err = v.registerCustomTagNameFunc()
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) registerCustomValidations() error {
	if vl, ok := v.engine.(*validator.Validate); ok {
		for _, f := range v.customValidationFuncs {
			err := vl.RegisterValidation(f.Tag, f.Func, f.CallValidationEvenIfNull)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("provided validation engine is not of type *validator.Validate")
}

// RegisterCustomValidation registers a custom tag name func to gin's validator
func (v *Validator) registerCustomTagNameFunc() error {
	if vl, ok := v.engine.(*validator.Validate); ok {
		vl.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		return nil
	}
	return errors.New("provided validation engine is not of type *validator.Validate")
}
