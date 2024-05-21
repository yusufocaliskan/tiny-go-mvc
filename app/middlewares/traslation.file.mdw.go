package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gptverse/init/framework/translator"
)

// Loads the translation file.
func LoadTranslationFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		selectedLang := ctx.GetHeader("Accept-Language")

		translation := translator.LoadErrorTextFile(selectedLang)
		ctx.Set("translations", translation)

		ctx.Next()
	}
}
