package v1routes

import (
	usercontroller "github.com/gptverse/init/app/controllers/users-controller"
	"github.com/gptverse/init/app/middlewares"
	authservice "github.com/gptverse/init/app/service/auth-service"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework"
)

func SetUserRoutes(fw *framework.Framework) {

	v1UserRoutes := fw.GinServer.Engine.Group("/api/v1/user")
	{
		//services
		uService := &userservice.UserService{Fw: fw, Collection: database.UserCollectionName}
		authService := &authservice.AuthService{Fw: fw, Collection: database.AuthCollectionName}

		//userControler
		uController := &usercontroller.UserController{Service: *uService, AuthService: *authService}

		//Creates new user
		v1UserRoutes.POST("/create/",
			middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&uController.User),
			middlewares.AuthCheck(fw),
			// middlewares.ForceOnlyRole("admin"),

			uController.Create)

		//Get by Id
		v1UserRoutes.GET("/fetch",
			middlewares.RateLimeter(),
			middlewares.AuthCheck(fw),
			middlewares.ValidateAndBind(&uController.UserWithIDFormIDModel),

			uController.Fetch)

		//Get by Id
		v1UserRoutes.GET("/fetch-all",
			middlewares.RateLimeter(),
			middlewares.AuthCheck(fw),
			middlewares.ValidateAndBind(&uController.UserFilterModel),

			uController.FetchAll)

		v1UserRoutes.PUT("/update",
			// middlewares.RateLimeter(),
			middlewares.AuthCheck(fw),
			middlewares.ValidateAndBind(&uController.UserUpdateModel),

			uController.Update)

		//Delete user
		//Only the one with {delete} permissions.
		v1UserRoutes.DELETE("/delete/",
			middlewares.RateLimeter(),
			middlewares.AuthCheck(fw),
			middlewares.ValidateAndBind(&uController.UserDeleteModel),

			uController.Delete)

	}
}
