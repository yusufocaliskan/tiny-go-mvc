package responser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gptverse/init/framework/translator"
)

/**
* Useage
* response.Code(http.StatusRequestTimeout).SetMessage(message).BadWithAbort()
* response.Payload(message).Success()
 */

type Response struct {
	Ctx        *gin.Context
	StatusCode int
	Data       interface{}
	Message    *translator.TranslationEntry
}

/* --------------------------------- setters -------------------------------- */

// status code
func (resp *Response) SetStatusCode(code int) *Response {
	resp.StatusCode = code
	return resp
}
func (resp *Response) SetMessage(text *translator.TranslationEntry) *Response {
	resp.Message = text
	return resp
}

// Sets the data
func (resp *Response) Payload(data interface{}) *Response {
	resp.Data = data
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
	data := gin.H{}

	if resp.Data != nil {
		data["data"] = resp.Data
	}

	data["message"] = resp.Message
	resp.CreateResponse(data, code)
}

// Bad request
func (resp *Response) Bad(err error) {

	data := gin.H{
		"message": resp.Message.Text,
	}
	if resp.Message.Code != "" {
		data = gin.H{
			"message": resp.Message.Text,
			"code":    resp.Message.Code,
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
		"message": resp.Message.Text,
	}
	if resp.Message.Code != "" {
		data = gin.H{
			"message": resp.Message.Text,
			"code":    resp.Message.Code,
		}
	}

	resp.Ctx.AbortWithStatusJSON(code, data)
}
