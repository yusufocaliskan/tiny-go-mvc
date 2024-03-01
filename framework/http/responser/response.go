package responser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tinyerror "github.com/yusufocaliskan/tiny-go-mvc/framework/http/tiny-error"
)

/**
* Useage
* response.Code(http.StatusRequestTimeout).SetError(message).BadWithAbort()
* response.Payload(message).Success()
 */

type Response struct {
	Ctx        *gin.Context
	StatusCode int
	Error      error
	Data       interface{}
}

/* --------------------------------- setters -------------------------------- */

// status code
func (resp *Response) Code(code int) *Response {
	resp.StatusCode = code
	return resp
}

// Sets the data
func (resp *Response) Payload(data interface{}) *Response {
	resp.Data = data
	return resp

}

// Set errror message
func (resp *Response) SetError(err string) *Response {
	resp.Error = tinyerror.New(err)
	return resp
}

/* ---------------------------------- resp ---------------------------------- */

func (resp *Response) CreateResponse(data interface{}, status int) {
	resp.Ctx.JSON(status, data)
}

// Return a success message with data
func (resp *Response) Success() {
	code := http.StatusOK
	if resp.StatusCode != 0 {
		code = resp.StatusCode
	}
	data := gin.H{
		"data": resp.Data,
		"code": code,
	}
	resp.CreateResponse(data, code)
}

// Bad request
func (resp *Response) Bad(err error) {

	data := gin.H{
		"error": resp.Error.Error(),
		"code":  http.StatusBadRequest,
	}
	resp.CreateResponse(data, http.StatusBadRequest)
}

// Stops the process, Use it when no need more execution
func (resp *Response) BadWithAbort() {
	code := http.StatusInternalServerError
	if resp.StatusCode != 0 {
		code = resp.StatusCode
	}
	data := gin.H{
		"error": resp.Error.Error(),
		"code":  code,
	}
	resp.Ctx.AbortWithStatusJSON(code, data)
}
