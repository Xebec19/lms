package requests

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
// 	re := regexp.MustCompile(`^[0-9]{10}$`) // basic 10-digit number
// 	return re.MatchString(fl.Field().String())
// })
