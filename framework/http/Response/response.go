package tinyresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func (resp *Response) Create(data interface{}, status int) {
	resp.Ctx.JSON(status, data)
}

// Return a success message with data
func (resp *Response) Success(data interface{}) {
	data = gin.H{
		"data": data,
		"code": http.StatusOK,
	}
	resp.Create(data, http.StatusOK)
}

// Bad request
func (resp *Response) Bad(err error) {
	data := gin.H{
		"error": err.Error(),
		"code":  http.StatusBadRequest,
	}
	resp.Create(data, http.StatusBadRequest)
}
