package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	usercontroller "github.com/gptverse/init/app/controllers/users-controller"
	"github.com/gptverse/init/config"
	"github.com/gptverse/init/framework"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
)

// Checking if the coming data valid
// AuthCheck validates the Authorization header token.
func AuthCheck(fw *framework.Framework, uController *usercontroller.UserController) gin.HandlerFunc {
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

		//Check Permissions
		CheckPermissions(uController, ctx, claims["data"].(string))

		ctx.Set("claim", claims["data"])
		ctx.Next()
	}
}

// Actions that would be allowed only for given role
func CheckPermissions(uController *usercontroller.UserController, ctx *gin.Context, emailAddress string) {
	// Get the role from the context
	role, exists := ctx.Get("UserRole")
	if !exists {
		err, user := uController.Service.GetUserByEmailAddress(emailAddress)
		if !err {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user data"})
			ctx.Abort()
			return
		}

		// Set the role in the context for future use
		role = user.Role
		ctx.Set("UserRole", user.Role)
	}

	requestMethod := strings.ToLower(ctx.Request.Method)

	requestPermission, ok := config.PermissionLookUp[requestMethod]
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Request method not allowed"})
		ctx.Abort()
		return
	}

	userRolePermissions, ok := config.DefinedPermissions[role.(string)]
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Role does not have defined permissions"})
		ctx.Abort()
		return
	}

	// Check if the user role has permission for the request method
	if !userRolePermissions[requestPermission] {
		// If not allowed, respond with an error
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Not allowed permission"})
		ctx.Abort()
		return
	}

	// Proceed to the next handler
	ctx.Next()
}
