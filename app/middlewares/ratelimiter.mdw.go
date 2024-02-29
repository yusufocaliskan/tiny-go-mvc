package middlewares

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RateLimeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ss := sessions.Default(ctx)

		var counter int
		if val, ok := ss.Get("RateLimiteCounter").(int); ok {

			counter = val
		}

		counter++
		ss.Set("RateLimiteCounter", counter)
		fmt.Println("RateLimiteCounter", counter)

		// ss.Delete("RateLimiteCounter")
		ss.Save()

		// Proceed with the request
		ctx.Next()
	}
}
