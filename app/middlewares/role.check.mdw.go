package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
)

// Check for the role
func RoleCheck(uController *usercontroller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		emailAddress, _ := ctx.Get("claim")
		_, getUserData := uController.Service.GetUserByEmailAddress(emailAddress.(string))

		requestMethod := strings.ToLower(ctx.Request.Method)
		userRolePermissions := config.DefinedPermissions[getUserData.Role]

		if !userRolePermissions[requestMethod] {

			ctx.JSON(http.StatusNotFound, gin.H{"error": "Not allowed permission"})
			ctx.Abort()
		}

		fmt.Println("Everything is okay")
		ctx.Next()
	}
}
