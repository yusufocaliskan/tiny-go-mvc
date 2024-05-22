package v1routes

import (
	usercontroller "github.com/gptverse/init/app/controllers/users-controller"
	"github.com/gptverse/init/app/middlewares"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/framework"
)

func SetUserRoutes(fw *framework.Framework) {

	v1UserRoutes := fw.GinServer.Engine.Group("/api/v1/user")
	{
		uService := &userservice.UserService{Fw: fw, Collection: "user"}
		uController := &usercontroller.UserController{Service: *uService}

		//Creates new user
		v1UserRoutes.POST("/createByEmail/",
			// middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&uController.User),
			middlewares.AuthCheck(fw, uController),
			// middlewares.ForceOnlyRole("admin"),

			uController.CreateNewUserByEmailAdress)

		//Get by Id
		v1UserRoutes.GET("/getUserInformationsById",
			middlewares.RateLimeter(),
			middlewares.AuthCheck(fw, uController),
			middlewares.ValidateAndBind(&uController.UserWithIDFormIDModel),

			uController.GetUserById)

		v1UserRoutes.PUT("/updateUserInformationsById",
			// middlewares.RateLimeter(),
			middlewares.AuthCheck(fw, uController),
			middlewares.ValidateAndBind(&uController.UserUpdateModel),

			uController.UpdateUserInformationsById)

		//Delete user
		//Only the one with {delete} permissions.
		v1UserRoutes.DELETE("/deleteById/",
			middlewares.RateLimeter(),
			middlewares.AuthCheck(fw, uController),
			middlewares.ValidateAndBind(&uController.UserDeleteModel),

			uController.DeleteUserById)

	}
}
