package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
)

func SetUserInformation2Session(fw *framework.Framework) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var secretKey = fw.Configs.AUTH_TOKEN_SECRET_KEY
		authHeader := ctx.GetHeader("Authorization")

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if the token is provided
		if bearerToken == "" {
			fmt.Println("Authorization header missing or empty")
			ctx.Next()
			return
		}

		uService := &userservice.UserService{Fw: fw, Collection: database.UserCollectionName}

		claims, tokenVerifyError := tinytoken.VerifyToken(bearerToken, secretKey)

		if tokenVerifyError != nil || claims == nil {
			fmt.Println("Invalid token or claims are nil")
			ctx.Next()
			return
		}

		emailAddress, ok := claims["data"].(string)

		if !ok || emailAddress == "" {
			fmt.Println("Email address claim is missing or empty")
			ctx.Next()
			return
		}

		sesStore := sessions.Default(ctx)
		oldUserInfos := sesStore.Get("CurrentUserInformations")

		if oldUserInfos == nil && emailAddress == "" {
			ctx.Next()
			return
		}

		// Fetch user information from the database
		userExists, user := uService.GetUserByEmailAddress(emailAddress)
		if !userExists {
			fmt.Println("User not found for email address:", emailAddress)
			ctx.Next()
			return
		}

		// Set the user information in the session
		sesStore.Set("CurrentUserInformations", user)
		sesStore.Set("BearerToken", bearerToken)
		if err := sesStore.Save(); err != nil {
			fmt.Println("Failed to save session:", err)
		}

		ctx.Next()
	}
}
