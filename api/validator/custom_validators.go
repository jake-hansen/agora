package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// meetingTimeValidateFunc checks a given string to ensure it is in appropriate RFC3339 format
// and ensures the time is in the future.
var meetingTimeValidateFunc validator.Func = func(fl validator.FieldLevel) bool {
	stringTime, ok := fl.Field().Interface().(string)
	if ok {
		t, err := time.Parse(time.RFC3339, stringTime)
		if err != nil {
			return false
		}
		now := time.Now().Add(time.Minute * -1)
		return !t.Before(now)
	}
	return false
}

// MeetingTimeValidator is a CustomValidationFunc that requires a string to be in RFC3339 format.
var MeetingTimeValidator = CustomValidationFunc{
	Tag:                      "valid meeting time",
	Func:                     meetingTimeValidateFunc,
	CallValidationEvenIfNull: false,
}
