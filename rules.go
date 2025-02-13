package validator

import (
	"regexp"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
)

func IsAlphaNumeric(str string, params ...string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z0-9]+([_:-]{1}[a-zA-Z0-9]+)*$", str)
	if !valid || err != nil {
		return false
	}
	if len(params) > 0 {
		if length, err := strconv.Atoi(params[0]); utf8.RuneCountInString(str) > length || err != nil {
			return false
		}
	}

	return true
}

func IsTag(str string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z]+([-]{1}[a-zA-Z]+)*$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 30 {
		return false
	}

	return true
}

func IsLanguage(str string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z]+([-]{1}[a-zA-Z]+)*$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}

	return true
}

func IsTopic(str string) bool {
	valid, err := regexp.MatchString("^[A-Z]+([_]{1}[A-Z]+)*$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}
	return true
}

func IsDisplayName(str string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z0-9]+(([a-zA-Z0-9 ])?[a-zA-Z0-9]*)*$", str)
	if !valid || err != nil {
		return false
	}

	if utf8.RuneCountInString(str) > 256 {
		return false
	}

	return true
}

func IsUserDisplayName(str string) bool {
	valid, err := regexp.MatchString("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}

	return true
}

func IsUUID4WithoutHyphens(str string) bool {
	valid, err := regexp.MatchString("^[0-9a-f]{16}[89ab][0-9a-f]{15}$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 32 {
		return false
	}
	return true
}

func IsOrderNumber(str string) bool {
	valid, err := regexp.MatchString("^O[0-9]{16}$", str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) != 17 {
		return false
	}

	return true
}

func IsDockerImage(str string) bool {
	valid, err := regexp.MatchString(`^[a-z0-9]+([\/_-]{1}[a-z0-9]+)*([:]{1}([0-9a-z]+([.-]{1}[0-9a-z]+)*)+)?$`, str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}
	return true
}

func IsOWASPEmail(str string) bool {
	valid, err := regexp.MatchString(`^[a-zA-Z0-9_+&*-]+(?:\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\.)+[a-zA-Z]{2,7}$`, str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 254 || utf8.RuneCountInString(str) < 5 {
		return false
	}
	return true
}

func IsPermissionResource(str string) bool {
	valid, err := regexp.MatchString(`^[A-Z]+([:]{1}([A-Z]+|(({[a-zA-Z]+})|[a-zA-Z0-9]+|\*)))*$`, str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}
	return true
}

func IsPath(str string) bool {
	valid, err := regexp.MatchString(`^(\/[a-zA-Z0-9]+)+$`, str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 256 {
		return false
	}
	return true
}

func IsURL(str string) bool {
	valid, err := regexp.MatchString(`^((((https?|ftps?|gopher|telnet|nntp):\/\/)|(mailto:|news:))(%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@&=+$,A-Za-z0-9])+)([).!';/?:,][[:blank:]])?$`, str) //nolint
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 2000 {
		return false
	}

	return true
}

func IsMemorySize(str string) bool {
	valid, err := regexp.MatchString(`^[0-9]+Mi$`, str)
	if !valid || err != nil {
		return false
	}
	if utf8.RuneCountInString(str) > 9 {
		return false
	}
	return true
}

func IsTime(str string) bool {
	return govalidator.IsTime(str, time.RFC3339)
}

func IsDate(str string) bool {
	valid, err := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, str)
	if !valid || err != nil {
		return false
	}
	return true
}

func IsJWT(str string) bool {
	valid, err := regexp.MatchString(`^([A-Za-z0-9\-_~+\/]+[=]{0,2})\.([A-Za-z0-9\-_~+\/]+[=]{0,2})(?:\.([A-Za-z0-9\-_~+\/]+[=]{0,2}))?$`, str) //nolint
	if !valid || err != nil {
		return false
	}
	return true
}

func IsNumeric(str string) bool {
	return govalidator.IsNumeric(str)
}

func IsIn(str string, params ...string) bool {
	return govalidator.IsIn(str, params...)
}

func IsLowerCase(str string) bool {
	return govalidator.IsLowerCase(str)
}

func IsPassword(str string) bool {
	fulfilledRules := 0
	rules := map[string]bool{
		"hasLengthValid": false,
		"hasUpper":       false,
		"hasLower":       false,
		"hasNumber":      false,
		"hasSpecial":     false,
	}

	if len(str) > 7 && len(str) <= 32 {
		rules["hasLengthValid"] = true
	}
	for _, char := range str {
		switch {
		case unicode.IsUpper(char):
			rules["hasUpper"] = true
		case unicode.IsLower(char):
			rules["hasLower"] = true
		case unicode.IsNumber(char):
			rules["hasNumber"] = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			rules["hasSpecial"] = true
		}
	}

	for rule, fulfilled := range rules {
		if fulfilled && rule != "hasLengthValid" {
			fulfilledRules++
		}
	}

	if rules["hasLengthValid"] && fulfilledRules >= 3 {
		return true
	}

	return false
}
