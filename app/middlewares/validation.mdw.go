package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	form "github.com/yusufocaliskan/tiny-go-mvc/framework/form/validate"
	responser "github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
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
			response.SetError(translator.GetMessage(ctx, bindingError.Error())).BadWithAbort()
		}

		fmt.Println("data--", data)
		//2. Check if is validated
		validationError, isError := validate.Check(data)

		if isError {
			// response.Error = validationError
			response.SetError(validationError).BadWithAbort()
		}

		ctx.Next()
	}
}
