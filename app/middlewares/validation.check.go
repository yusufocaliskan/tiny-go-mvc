package middlewares

import (
	"github.com/gin-gonic/gin"
	form "github.com/yusufocaliskan/tiny-go-mvc/framework/form/validate"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/Response"
)

type ValidationCheck struct{}

// Checking if the coming database valid
func (vCheck ValidationCheck) IsValidate(data interface{}) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		validate := form.FormValidator{}
		response := tinyresponse.Response{Ctx: ctx}

		ctx.BindJSON(&data)

		//Check if is validated
		validationError := validate.Check(data)

		if validationError != nil {
			response.Bad(validationError)
			ctx.Abort()
		}

		ctx.Next()
	}
}
