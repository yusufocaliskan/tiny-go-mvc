package v1routes

import (
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/app/middlewares"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

func SetUserRoutes(fw *framework.Framework) {
	v1UserRoutes := fw.GinServer.Engine.Group("/api/v1/user")
	{
		uService := &userservice.UserService{Fw: fw, Collection: "user"}
		uController := &usercontroller.UserController{Service: *uService}

		//Creates new user
		v1UserRoutes.POST("/createByEmail/",

			middlewares.AuthCheck(fw),
			middlewares.Check4ValidData(&uController.User),
			middlewares.RateLimeter(),
			uController.CreateNewUserByEmailAdress)

	}
}
