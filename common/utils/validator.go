package utils

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}

// GetValidateStruct validates a struct using the validator package
func GetValidatedStruct(r *http.Request, s interface{}) error {

	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return err
	}

	validate := GetValidator()
	err = validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
