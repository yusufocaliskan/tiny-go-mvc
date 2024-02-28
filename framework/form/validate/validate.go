package form

import "github.com/go-playground/validator/v10"

var validatorInstance = validator.New()

type FormValidator struct{}

func (Fv FormValidator) Check(s interface{}) (validationError error) {

	validationError = validatorInstance.Struct(s)

	return

}
