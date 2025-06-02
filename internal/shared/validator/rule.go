package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-z0-9_]{3,20}$`)
	phoneRegex    = regexp.MustCompile(`^0\d{9}$`)
)

func UsernameValid(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

func PhoneVN(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(fl.Field().String())
}

func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber bool
	for _, c := range password {
		switch {
		case 'a' <= c && c <= 'z':
			hasLower = true
		case 'A' <= c && c <= 'Z':
			hasUpper = true
		case '0' <= c && c <= '9':
			hasNumber = true
		}
	}

	return hasUpper && hasLower && hasNumber
}
