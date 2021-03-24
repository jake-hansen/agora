package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationEngine interface{}

// CustomValidationFunc represents a custom validation function and other associated properties.
type CustomValidationFunc struct {
	Tag                      string
	Func                     validator.Func
	CallValidationEvenIfNull bool
}

// Config is used to configure a Validator with the containing Engine and CustomValidationFuncs.
type Config struct {
	Engine                ValidationEngine
	CustomValidationFuncs []CustomValidationFunc
}

// Validator contains the components needed to perform validation within the application.
type Validator struct {
	engine                ValidationEngine
	customValidationFuncs []CustomValidationFunc
}

// NewValidator returns a new Validator configured using the provided Config and also
// initializes the Validator for use.
func NewValidator(config Config) (*Validator, error) {
	v := &Validator{
		engine:                config.Engine,
		customValidationFuncs: config.CustomValidationFuncs,
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
	err = v.registerCustomValidations()
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
