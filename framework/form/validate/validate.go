package form

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
)

var validatorInstance = validator.New()
var (
	uni *ut.UniversalTranslator
)

type FormValidator struct{}

func (Fv FormValidator) Check(s interface{}) (*translator.TranslationEntry, bool) {

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	if err := en_translations.RegisterDefaultTranslations(validatorInstance, trans); err != nil {
		panic("failed to register translations: " + err.Error())
	}

	validationError := validatorInstance.Struct(s)
	errs, isError := TranslateError(validationError, trans)
	if isError {

		return &translator.TranslationEntry{Text: errs, Code: "UF000"}, true
	}

	return &translator.TranslationEntry{Text: "", Code: ""}, false

}

func TranslateError(err error, trans ut.Translator) (string, bool) {
	if err == nil {
		return "", false
	}
	validatorErrs := err.(validator.ValidationErrors)
	var result string

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		result = translatedErr.Error()
	}

	return result, true
}
