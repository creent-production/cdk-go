package validation

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	phoneIdRegexString        = "^(\\+62|62|0)8[1-9][0-9]{6,9}$"
	pathFileStaticRegexString = "^.*static.*$"
)

var (
	phoneIdRegex        = regexp.MustCompile(phoneIdRegexString)
	pathFileStaticRegex = regexp.MustCompile(pathFileStaticRegexString)
)

func validatePathFileStatic(fl validator.FieldLevel) bool {
	return pathFileStaticRegex.MatchString(fl.Field().String())
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	switch fl.Param() {
	case "id":
		return phoneIdRegex.MatchString(fl.Field().String())
	default:
		return false
	}
}

func containsAlpha(s string) bool {
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	for _, char := range s {
		if strings.Contains(alpha, strings.ToLower(string(char))) {
			return true
		}
	}
	return false
}

func countDigit(s string) int {
	const digit = "0123456789"
	count := 0
	for _, char := range s {
		if strings.Contains(digit, strings.ToLower(string(char))) {
			count += 1
		}
	}
	return count
}

func validateNoIdentity(fl validator.FieldLevel) bool {
	identity := fl.Field().String()
	digitSum := countDigit(identity)
	if containsAlpha(identity) && digitSum < 10 {
		return len(identity) > 5 && len(identity) < 15 && digitSum >= 4
	}
	return digitSum == 16
}

func validateDateTimeRange(fl validator.FieldLevel, kind string) bool {
	param := strings.Split(fl.Param(), "-")
	if len(param) != 3 {
		return false
	}

	var timezone string
	switch param[0] {
	case "wib":
		timezone = "Asia/Jakarta"
	case "wita":
		timezone = "Asia/Ujung_Pandang"
	case "wit":
		timezone = "Asia/Jayapura"
	default:
		return false
	}

	//init the loc
	loc, _ := time.LoadLocation(timezone)

	layoutFormat := "2006-01-02 15:04:05"
	if kind == "date" {
		layoutFormat = "2006-01-02"
	}
	date, err := time.ParseInLocation(layoutFormat, fl.Field().String(), loc)
	if err != nil {
		return false
	}

	//set timezone
	now := time.Now().In(loc)
	if kind == "date" {
		now, _ = time.Parse(layoutFormat, now.Format(layoutFormat))
		date, _ = time.Parse(layoutFormat, date.Format(layoutFormat))
	}

	durBefore, err := time.ParseDuration(param[1])
	if err != nil {
		return false
	}

	durAfter, err := time.ParseDuration(param[2])
	if err != nil {
		return false
	}

	return now.Add(-durBefore).Unix() <= date.Unix() && now.Add(durAfter).Unix() >= date.Unix()
}

func validateDateTimeDuration(fl validator.FieldLevel, kind string) bool {
	param := strings.Split(fl.Param(), "-")
	if len(param) != 3 {
		return false
	}

	var timezone string
	switch param[0] {
	case "wib":
		timezone = "Asia/Jakarta"
	case "wita":
		timezone = "Asia/Ujung_Pandang"
	case "wit":
		timezone = "Asia/Jayapura"
	default:
		return false
	}

	//init the loc
	loc, _ := time.LoadLocation(timezone)

	layoutFormat := "2006-01-02 15:04:05"
	if kind == "date" {
		layoutFormat = "2006-01-02"
	}
	date, err := time.ParseInLocation(layoutFormat, fl.Field().String(), loc)
	if err != nil {
		return false
	}

	//set timezone
	now := time.Now().In(loc)
	if kind == "date" {
		now, _ = time.Parse(layoutFormat, now.Format(layoutFormat))
		date, _ = time.Parse(layoutFormat, date.Format(layoutFormat))
	}

	dur, err := time.ParseDuration(param[2])
	if err != nil {
		return false
	}

	switch param[1] {
	case "lt":
		return date.Unix() < now.Add(-dur).Unix()
	case "lte":
		return date.Unix() <= now.Add(-dur).Unix()
	case "gt":
		return date.Unix() > now.Add(dur).Unix()
	case "gte":
		return date.Unix() >= now.Add(dur).Unix()
	default:
		return false
	}
}

func validateDateTime(fl validator.FieldLevel, kind string) bool {
	param := strings.Split(fl.Param(), "-")
	if len(param) != 2 {
		return false
	}

	var timezone string
	switch param[0] {
	case "wib":
		timezone = "Asia/Jakarta"
	case "wita":
		timezone = "Asia/Ujung_Pandang"
	case "wit":
		timezone = "Asia/Jayapura"
	default:
		return false
	}

	//init the loc
	loc, _ := time.LoadLocation(timezone)

	layoutFormat := "2006-01-02 15:04:05"
	if kind == "date" {
		layoutFormat = "2006-01-02"
	}
	date, err := time.ParseInLocation(layoutFormat, fl.Field().String(), loc)
	if err != nil {
		return false
	}

	//set timezone
	now := time.Now().In(loc)
	if kind == "date" {
		now, _ = time.Parse(layoutFormat, now.Format(layoutFormat))
		date, _ = time.Parse(layoutFormat, date.Format(layoutFormat))
	}

	switch param[1] {
	case "lt":
		return date.Unix() < now.Unix()
	case "lte":
		return date.Unix() <= now.Unix()
	case "gt":
		return date.Unix() > now.Unix()
	case "gte":
		return date.Unix() >= now.Unix()
	case "none":
		return true
	default:
		return false
	}
}

func validateDateTimeField(fl validator.FieldLevel, kind string) bool {
	if len(fl.Field().String()) < 1 {
		return true
	}

	param := strings.Split(fl.Param(), "-")
	if len(param) != 2 {
		return false
	}

	field := fl.Top().FieldByName(param[1])
	if !field.IsValid() {
		return false
	}

	layoutFormat := "2006-01-02 15:04:05"
	if kind == "date" {
		layoutFormat = "2006-01-02"
	}

	currentDate, err := time.Parse(layoutFormat, fl.Field().String())
	if err != nil {
		return false
	}

	date, err := time.Parse(layoutFormat, field.String())
	if err != nil {
		return false
	}

	switch param[0] {
	case "lt":
		return currentDate.Unix() < date.Unix()
	case "lte":
		return currentDate.Unix() <= date.Unix()
	case "gt":
		return currentDate.Unix() > date.Unix()
	case "gte":
		return currentDate.Unix() >= date.Unix()
	default:
		return false
	}
}

func validateDateZoneDuration(fl validator.FieldLevel) bool {
	return validateDateTimeDuration(fl, "date")
}

func validateDateTimeZoneDuration(fl validator.FieldLevel) bool {
	return validateDateTimeDuration(fl, "datetime")
}

func validateDateZoneRange(fl validator.FieldLevel) bool {
	return validateDateTimeRange(fl, "date")
}

func validateDateTimeZoneRange(fl validator.FieldLevel) bool {
	return validateDateTimeRange(fl, "datetime")
}

func validateDateZone(fl validator.FieldLevel) bool {
	return validateDateTime(fl, "date")
}

func validateDateTimeZone(fl validator.FieldLevel) bool {
	return validateDateTime(fl, "datetime")
}

func validateDateZoneField(fl validator.FieldLevel) bool {
	return validateDateTimeField(fl, "date")
}

func validateDateTimeZoneField(fl validator.FieldLevel) bool {
	return validateDateTimeField(fl, "datetime")
}

func validateTime(fl validator.FieldLevel) bool {
	layoutFormat := "15:04:05"

	_, err := time.Parse(layoutFormat, fl.Field().String())
	if err != nil {
		return false
	}

	return true
}

func validateTimeField(fl validator.FieldLevel) bool {
	if len(fl.Field().String()) < 1 {
		return true
	}

	param := strings.Split(fl.Param(), "-")
	if len(param) != 2 {
		return false
	}

	field := fl.Top().FieldByName(param[1])
	if !field.IsValid() {
		return false
	}

	layoutFormat := "15:04:05"

	currentTime, err := time.Parse(layoutFormat, fl.Field().String())
	if err != nil {
		return false
	}

	times, err := time.Parse(layoutFormat, field.String())
	if err != nil {
		return false
	}

	switch param[0] {
	case "lt":
		return currentTime.Unix() < times.Unix()
	case "lte":
		return currentTime.Unix() <= times.Unix()
	case "gt":
		return currentTime.Unix() > times.Unix()
	case "gte":
		return currentTime.Unix() >= times.Unix()
	default:
		return false
	}
}

func registerValidation(validate *validator.Validate) {
	validate.RegisterValidation("time", validateTime)
	validate.RegisterValidation("datezone", validateDateZone)
	validate.RegisterValidation("datetimezone", validateDateTimeZone)
	validate.RegisterValidation("datezoneduration", validateDateZoneDuration)
	validate.RegisterValidation("datetimezoneduration", validateDateTimeZoneDuration)
	validate.RegisterValidation("datezonerange", validateDateZoneRange)
	validate.RegisterValidation("datetimezonerange", validateDateTimeZoneRange)
	validate.RegisterValidation("timefield", validateTimeField)
	validate.RegisterValidation("datefield", validateDateZoneField)
	validate.RegisterValidation("datetimefield", validateDateTimeZoneField)
	validate.RegisterValidation("phone", validatePhoneNumber)
	validate.RegisterValidation("noidentity", validateNoIdentity)
	validate.RegisterValidation("pathfilestatic", validatePathFileStatic)
}
