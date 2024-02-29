package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	tinysession "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-session"
)

func RateLimeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("Hereeeee worksssss....")

		tinySession := &tinysession.TinySession{}
		session := tinySession.New(ctx)

		session.Set("test", "test session value")
		fmt.Println("Valuessssss===", session.Get("test"))
		session.Save()
		// Proceed with the request
		ctx.Next()
	}
}
