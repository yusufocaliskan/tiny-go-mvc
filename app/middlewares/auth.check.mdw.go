package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
)

// Checking if the coming data valid
// AuthCheck validates the Authorization header token.
func AuthCheck(allowedRole string, fw *framework.Framework, uController *usercontroller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var secretKey = fw.Configs.AUTH_TOKEN_SECRET_KEY
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
		claims, err := tinytoken.VerifyToken(bearerToken, secretKey)
		if err != nil {
			fmt.Println("Error: ", err)
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}
		CheckPermissions(allowedRole, uController, ctx, claims["data"].(string))

		fmt.Println("Auth Check successful, claims:", claims["data"])
		ctx.Set("claim", claims["data"])
		ctx.Next()
	}
}

func CheckPermissions(allowedRole string, uController *usercontroller.UserController, ctx *gin.Context, emailAddress string) {
	_, getUserData := uController.Service.GetUserByEmailAddress(emailAddress)

	requestMethod := strings.ToLower(ctx.Request.Method)
	userRolePermissions := config.DefinedPermissions[getUserData.Role]

	if allowedRole != getUserData.Role {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not allowed permission"})
		ctx.Abort()
	}

	//is user allowed for the method?
	if !userRolePermissions[requestMethod] {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not allowed permission"})
		ctx.Abort()
	}

	fmt.Println("Everything is okay")
	ctx.Next()
}
