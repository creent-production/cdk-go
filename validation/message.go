package validation

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	Default      = "Invalid value."
	Required     = "Missing data for required field."
	Email        = "Not a valid email address."
	Min          = "Shorter than minimum length %s."
	Max          = "Longer than maximum length %s."
	Len          = "Length must be equal to %s."
	Gte          = "Must be greater than or equal to %s."
	Gt           = "Must be greater than %s."
	Lte          = "Must be less than or equal to %s."
	Lt           = "Must be less than %s."
	Oneof        = "Must be one of: %s."
	Eqfield      = "Must be equal to %s."
	Unique       = "Must be unique."
	Phone        = "Invalid phone number."
	DateZone     = "Invalid date."
	DateTimeZone = "Invalid date time."
	NoIdentity   = "Invalid NIK/NIORA."
)

func SetErrorMessage(ev ErrorValidate) string {
	switch ev.Tag {
	case "required":
		return Required
	case "email":
		return Email
	case "min":
		return fmt.Sprintf(Min, ev.Param)
	case "max":
		return fmt.Sprintf(Max, ev.Param)
	case "len":
		return fmt.Sprintf(Len, ev.Param)
	case "gte":
		msg := ev.Param
		if len(msg) == 0 {
			msg = "current time"
		}
		return fmt.Sprintf(Gte, msg)
	case "gt":
		msg := ev.Param
		if len(msg) == 0 {
			msg = "current time"
		}
		return fmt.Sprintf(Gt, msg)
	case "lte":
		msg := ev.Param
		if len(msg) == 0 {
			msg = "current time"
		}
		return fmt.Sprintf(Lte, msg)
	case "lt":
		msg := ev.Param
		if len(msg) == 0 {
			msg = "current time"
		}
		return fmt.Sprintf(Lt, msg)
	case "oneof":
		var msg string
		reg := regexp.MustCompile(`'(.*?)'`)
		txt := reg.FindAllString(ev.Param, -1)
		if len(txt) == 0 {
			msg = strings.Join(strings.Split(ev.Param, " "), ", ")
		} else {
			msg = strings.Join(txt, ", ")
		}
		return fmt.Sprintf(Oneof, msg)
	case "eqfield":
		return fmt.Sprintf(Eqfield, ToSnakeCase(ev.Param))
	case "unique":
		return Unique
	case "phone":
		return Phone
	case "datezone":
		return DateZone
	case "datetimezone":
		return DateTimeZone
	case "noidentity":
		return NoIdentity
	}
	return Default
}
