package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	form "github.com/gptverse/init/framework/form/validate"
	responser "github.com/gptverse/init/framework/http/responser"
	"github.com/gptverse/init/framework/translator"
)

// Checking if the coming data valid
func Check4ValidData(data interface{}) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		validate := form.FormValidator{}
		response := responser.Response{Ctx: ctx}

		//1. Binding the incoming data with the struct
		bindingError := ctx.BindJSON(&data)

		if bindingError != nil {

			print("bindingError", bindingError.Error())
			// response.Error = bindingError
			response.SetMessage(translator.GetMessage(ctx, bindingError.Error())).BadWithAbort()
		}

		fmt.Println("data--", data)
		//2. Check if is validated
		validationError, isError := validate.Check(data)

		if isError {
			// response.Error = validationError
			response.SetMessage(validationError).BadWithAbort()
		}

		ctx.Next()
	}
}
