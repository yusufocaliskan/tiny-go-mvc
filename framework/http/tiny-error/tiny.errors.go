package tinyerror

import "github.com/yusufocaliskan/tiny-go-mvc/framework/translator"

// Create a custom error message
func New(err *translator.TranslationEntry) *translator.TranslationEntry {
	return err
}
