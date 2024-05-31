package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	form "github.com/yusufocaliskan/tiny-go-mvc/framework/form/validate"
	responser "github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
)

const (
	contentTypeMultipartForm  = "multipart/form-data"
	contentTypeJSON           = "application/json"
	contentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

// ValidateAndBind checks if the incoming data is valid and binds it to the provided struct
func ValidateAndBind(data interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validate := form.FormValidator{}
		response := responser.Response{Ctx: ctx}

		var bindingError error

		switch ctx.ContentType() {

		//File
		case contentTypeMultipartForm:
			bindingError = ctx.ShouldBind(data)

		// GET
		case contentTypeFormURLEncoded:
			bindingError = ctx.ShouldBind(data)

		// POST JSON
		case contentTypeJSON:
			bindingError = ctx.ShouldBindJSON(data)

		default:
			if ctx.Request.Method == http.MethodGet {
				bindingError = ctx.ShouldBindQuery(data)
			} else {
				bindingError = ctx.ShouldBindJSON(data)
			}
		}

		//There should'nt any error
		if bindingError != nil {
			response.SetMessage(translator.GetMessage(ctx, bindingError.Error())).BadWithAbort()
			return
		}

		//Validate
		if validationError, isError := validate.Check(data); isError {
			response.SetMessage(validationError).BadWithAbort()
			return
		}

		ctx.Next()
	}
}
