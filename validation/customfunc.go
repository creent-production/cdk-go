package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const (
	phoneIdRegexString = "^(\\+62|62|0)8[1-9][0-9]{6,9}$"
)

var (
	phoneIdRegex = regexp.MustCompile(phoneIdRegexString)
)

func validatePhoneNumber(fl validator.FieldLevel) bool {
	switch fl.Param() {
	case "id":
		return phoneIdRegex.MatchString(fl.Field().String())
	default:
		return false
	}
}

func registerValidation(validate *validator.Validate) {
	validate.RegisterValidation("phone", validatePhoneNumber)
}
