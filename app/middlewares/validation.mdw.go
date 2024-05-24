package middlewares

import (
	"github.com/gin-gonic/gin"
	form "github.com/gptverse/init/framework/form/validate"
	responser "github.com/gptverse/init/framework/http/responser"
	"github.com/gptverse/init/framework/translator"
)

// Checking if the coming data valid
func ValidateAndBind(data interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validate := form.FormValidator{}
		response := responser.Response{Ctx: ctx}

		var bindingError error

		switch ctx.Request.Method {
		case "GET":
			bindingError = ctx.ShouldBindQuery(data)
		default:
			bindingError = ctx.ShouldBindJSON(data)
		}

		if bindingError != nil {
			response.SetMessage(translator.GetMessage(ctx, bindingError.Error())).BadWithAbort()
			return
		}

		if validationError, isError := validate.Check(data); isError {
			response.SetMessage(validationError).BadWithAbort()
		}

		ctx.Next()
	}
}
