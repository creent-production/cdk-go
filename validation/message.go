package validation

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	Default           = "Invalid value."
	Required          = "Missing data for required field."
	Email             = "Not a valid email address."
	Min               = "Shorter than minimum length %s."
	Max               = "Longer than maximum length %s."
	Len               = "Length must be equal to %s."
	Gte               = "Must be greater than or equal to %s."
	Gt                = "Must be greater than %s."
	Lte               = "Must be less than or equal to %s."
	Lt                = "Must be less than %s."
	Oneof             = "Must be one of: %s."
	Eqfield           = "Must be equal to %s."
	Unique            = "Must be unique."
	Phone             = "Invalid phone number."
	DateZone          = "Invalid date."
	DateTimeZone      = "Invalid date time."
	DateTimeZoneRange = "Must be between range (now - %s) until (now + %s)."
	Time              = "Invalid time."
	Number            = "Invalid number."
	Numeric           = "Invalid numeric."
	NoIdentity        = "Invalid NIK/NIORA."
)

func SetErrorMessage(ev ErrorValidate) string {
	switch ev.Tag {
	case "datezonerange", "datetimezonerange":
		param := strings.Split(ev.Param, "-")
		return fmt.Sprintf(DateTimeZoneRange, param[1], param[2])
	case "required", "required_with", "required_if":
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
		param := strings.Split(ev.Param, "-")
		switch param[1] {
		case "lt":
			return fmt.Sprintf(Lt, "date now")
		case "lte":
			return fmt.Sprintf(Lte, "date now")
		case "gt":
			return fmt.Sprintf(Gt, "date now")
		case "gte":
			return fmt.Sprintf(Gte, "date now")
		}
		return DateZone
	case "datetimezone":
		param := strings.Split(ev.Param, "-")
		switch param[1] {
		case "lt":
			return fmt.Sprintf(Lt, "date time now")
		case "lte":
			return fmt.Sprintf(Lte, "date time now")
		case "gt":
			return fmt.Sprintf(Gt, "date time now")
		case "gte":
			return fmt.Sprintf(Gte, "date time now")
		}
		return DateTimeZone
	case "datefield":
		param := strings.Split(ev.Param, "-")
		switch param[0] {
		case "lt":
			return fmt.Sprintf(Lt, ToSnakeCase(param[1]))
		case "lte":
			return fmt.Sprintf(Lte, ToSnakeCase(param[1]))
		case "gt":
			return fmt.Sprintf(Gt, ToSnakeCase(param[1]))
		case "gte":
			return fmt.Sprintf(Gte, ToSnakeCase(param[1]))
		}
		return DateZone
	case "datetimefield":
		param := strings.Split(ev.Param, "-")
		switch param[0] {
		case "lt":
			return fmt.Sprintf(Lt, ToSnakeCase(param[1]))
		case "lte":
			return fmt.Sprintf(Lte, ToSnakeCase(param[1]))
		case "gt":
			return fmt.Sprintf(Gt, ToSnakeCase(param[1]))
		case "gte":
			return fmt.Sprintf(Gte, ToSnakeCase(param[1]))
		}
		return DateTimeZone
	case "timefield":
		param := strings.Split(ev.Param, "-")
		switch param[0] {
		case "lt":
			return fmt.Sprintf(Lt, ToSnakeCase(param[1]))
		case "lte":
			return fmt.Sprintf(Lte, ToSnakeCase(param[1]))
		case "gt":
			return fmt.Sprintf(Gt, ToSnakeCase(param[1]))
		case "gte":
			return fmt.Sprintf(Gte, ToSnakeCase(param[1]))
		}
		return Time
	case "time":
		return Time
	case "noidentity":
		return NoIdentity
	case "numeric":
		return Numeric
	case "number":
		return Number
	}
	return Default
}
