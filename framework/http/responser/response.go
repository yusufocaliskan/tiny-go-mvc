package responser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
)

/**
* Useage
* response.Code(http.StatusRequestTimeout).SetError(message).BadWithAbort()
* response.Payload(message).Success()
 */

type Response struct {
	Ctx        *gin.Context
	StatusCode int
	Error      *translator.TranslationEntry
	Data       interface{}
	Message    string
}

/* --------------------------------- setters -------------------------------- */

// status code
func (resp *Response) SetStatusCode(code int) *Response {
	resp.StatusCode = code
	return resp
}
func (resp *Response) SetMessage(text string) *Response {
	resp.Message = text
	return resp
}

// Sets the data
func (resp *Response) Payload(data interface{}) *Response {
	resp.Data = data
	return resp

}

// Set errror message
func (resp *Response) SetError(err *translator.TranslationEntry) *Response {
	resp.Error = err
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
		"code": code,
	}

	if resp.Data != nil {
		data["data"] = resp.Data
	}

	if resp.Message != "" {
		data["Message"] = resp.Message
	}

	resp.CreateResponse(data, code)
}

// Bad request
func (resp *Response) Bad(err error) {

	data := gin.H{
		"error": resp.Error.Text,
	}
	if resp.Error.Code != "" {
		data = gin.H{
			"error": resp.Error.Text,
			"code":  resp.Error.Code,
		}
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
		"error": resp.Error.Text,
	}
	if resp.Error.Code != "" {
		data = gin.H{
			"error": resp.Error.Text,
			"code":  resp.Error.Code,
		}
	}

	resp.Ctx.AbortWithStatusJSON(code, data)
}
