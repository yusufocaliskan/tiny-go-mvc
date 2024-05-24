package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	usermodel "github.com/gptverse/init/app/models/user-model"
	"github.com/gptverse/init/config"
	"github.com/gptverse/init/framework"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
)

// Checking if the coming data valid
// AuthCheck validates the Authorization header token.
func AuthCheck(fw *framework.Framework) gin.HandlerFunc {
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
		CheckPermissions(fw, ctx, claims["data"].(string))

		ctx.Set("claim", claims["data"])
		ctx.Next()
	}
}

// Actions that would be allowed only for given role
func CheckPermissions(fw *framework.Framework, ctx *gin.Context, emailAddress string) {

	// Get the role from the context
	sesStore := sessions.Default(ctx)
	fetchCurrentUserInfo := sesStore.Get("CurrentUserInformations")

	currentUserInfo, ok := fetchCurrentUserInfo.(*usermodel.UserModel)
	if !ok {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user data"})
		ctx.Abort()
		return
	}

	if currentUserInfo.Role == "" {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user data"})
		ctx.Abort()
		return

	}

	requestMethod := strings.ToLower(ctx.Request.Method)
	requestPermission, ok := config.PermissionLookUp[requestMethod]

	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Request method not allowed"})
		ctx.Abort()
		return
	}

	userRolePermissions, ok := config.DefinedPermissions[currentUserInfo.Role]

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
