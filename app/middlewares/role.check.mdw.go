package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
)

// Check for the role
func RoleCheck(uController *usercontroller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err, getUserData := uController.Service.GetUserById(uController.UserDeleteModel.Id)

		if err {

			fmt.Println("Uesr Not found", uController.UserDeleteModel.Id)
			return
		}

		fmt.Println("uController.UserDeleteModel", uController.UserDeleteModel.Id, getUserData)
		ctx.Next()
	}
}
