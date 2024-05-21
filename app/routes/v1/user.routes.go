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

			middlewares.RateLimeter(),
			middlewares.Check4ValidData(&uController.User),
			middlewares.AuthCheck(fw, uController),

			// middlewares.RateLimeter(),
			uController.CreateNewUserByEmailAdress)

		//Delete user
		v1UserRoutes.DELETE("/deleteById/",
			middlewares.RateLimeter(),

			middlewares.AuthCheck(fw, uController),
			middlewares.Check4ValidData(&uController.UserDeleteModel),
			middlewares.ForceOnlyRole("admin"),

			uController.DeleteUserById)

	}
}
