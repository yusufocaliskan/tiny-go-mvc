package tinyerror

import "github.com/gptverse/init/framework/translator"

// Create a custom error message
func New(err *translator.TranslationEntry) *translator.TranslationEntry {
	return err
}
