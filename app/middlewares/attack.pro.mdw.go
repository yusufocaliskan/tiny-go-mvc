package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Proctection.
func AttackProtectionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

        ctx.Header("X-Content-Type-Options", "nosniff")
        ctx.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")

		ctx.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		ctx.Writer.Header().Set("X-Frame-Options", "DENY")
		ctx.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Next()
	}
}
