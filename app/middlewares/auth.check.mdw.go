package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	authservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/auth-service"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
)

// Checking if the coming data valid
// AuthCheck validates the Authorization header token.
func AuthCheck(fw *framework.Framework, authService *authservice.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var secretKey = fw.Configs.AUTH_TOKEN_SECRET_KEY
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
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

		//Fetch user token from database
		///todo: cache it
		emailAddress := claims["data"].(string)
		_, tokenInDatabase := authService.GetToken(emailAddress)
		if tokenInDatabase != nil {

			if tokenInDatabase.Status == "passive" {

				ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
				return
			}

			if isExpired(time.Now(), tokenInDatabase.Token.AccessToken.ExpiryTime) {

				ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			}

		}

		//Check Permissions
		CheckPermissions(fw, ctx, emailAddress)

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

func isExpired(issueTime time.Time, duration time.Duration) bool {
	expirationTime := issueTime.Add(duration)
	return expirationTime.Before(time.Now())
}
