package validator

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NewNoSpecialCharacterCustomValidator() validator.Func {
	allowedChars := []string{
		"\u0030-\u0039",              // 0-9
		"\u0041-\u005A",              // A-Z
		"\u0061-\u007A",              // a-z
		"\u0E01-\u0E3A\u0E40-\u0E4E", // Thai letters
		"\u0028",                     // (
		"\u0029",                     // )
		"\u002C",                     // ,
		"\u002F",                     // /
		"\u201C",                     // “ Left Double Quotation Mark
		"\u201D",                     // ” Right Double Quotation Mark
		"\u0022",                     // " Quotation Mark
		"\u003A",                     // :
		"\u0040",                     // @
		"\u005F",                     // _
		"\u0020",                     // space
		"\u0026",                     // &
		"\u002E",                     // .
		"\u2013",                     // – En Dash
		"\u002D",                     // - Hyphen-Minus
	}

	notAllowedChars := []string{
		"\u2013\u2013", // double – En Dash
		"\u002D\u002D", // double - Hyphen-Minus
	}

	pattern := fmt.Sprintf("([^%s]|%s)+", strings.Join(allowedChars, ""), strings.Join(notAllowedChars, "|"))

	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	return func(fl validator.FieldLevel) bool {
		switch field := fl.Field().Interface().(type) {
		case string:
			if field == "" {
				return true
			}

			if r.MatchString(field) {
				return false
			}
		}

		return true
	}
}

// validate only number size
// don't validate decimal size
// must use only float64
// will failed on float32
func NewDecimalCustomValidator() validator.Func {
	return func(fl validator.FieldLevel) bool {
		switch value := fl.Field().Interface().(type) {
		case float32:
			return false
		case float64:
			if value == 0 {
				return true
			}

			if value < 0 {
				return false
			}

			params := strings.Split(fl.Param(), "-")

			if len(params) != 2 {
				return false
			}

			numberSize, err := strconv.Atoi(params[0])
			if err != nil {
				return false
			}

			decimalSize, err := strconv.Atoi(params[1])
			if err != nil {
				return false
			}

			if value >= math.Pow10(numberSize-decimalSize) {
				return false
			}

			return true
		default:
			return true
		}
	}
}
