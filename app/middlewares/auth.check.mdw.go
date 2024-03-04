package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
)

// Checking if the coming data valid
// AuthCheck validates the Authorization header token.
func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		// Ensure the Authorization header is formatted as "Bearer <token>"
		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if bearerToken == authHeader {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		// Verify the token
		claims, err := tinytoken.VerifyToken(bearerToken)
		if err != nil {
			fmt.Println("Error: ", err)
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		fmt.Println("Auth Check successful, claims:", claims["data"])
		ctx.Next()
	}
}
