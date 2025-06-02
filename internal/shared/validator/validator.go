package validator

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	// Register các custom rule
	Validate.RegisterValidation("username_valid", UsernameValid)
	Validate.RegisterValidation("phone_vn", PhoneVN)
	Validate.RegisterValidation("password_strong", StrongPassword)
}
