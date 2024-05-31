package v1routes

import (
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/app/middlewares"
	authservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/auth-service"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
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
			middlewares.ValidateAndBind(&uController.User),
			middlewares.AuthCheck(fw, authService),
			// middlewares.ForceOnlyRole("admin"),

			uController.Create)

		//Get by Id
		v1UserRoutes.GET("/fetch",
			middlewares.AuthCheck(fw, authService),
			middlewares.ValidateAndBind(&uController.UserWithIDFormIDModel),

			uController.Fetch)

		//Get by Id
		v1UserRoutes.GET("/fetch-all",
			middlewares.AuthCheck(fw, authService),
			middlewares.ValidateAndBind(&uController.UserFilterModel),

			uController.FetchAll)

		v1UserRoutes.PUT("/update",
			middlewares.AuthCheck(fw, authService),
			middlewares.ValidateAndBind(&uController.UserUpdateModel),

			uController.Update)

		//Delete user
		//Only the one with {delete} permissions.
		v1UserRoutes.DELETE("/delete/",
			middlewares.AuthCheck(fw, authService),
			middlewares.ValidateAndBind(&uController.UserDeleteModel),

			uController.Delete)

	}
}
