package v1routes

import (
	authcontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/auth-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/app/middlewares"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

func SetAuthRoutes(fw *framework.Framework) {
	v1AuthRoutes := fw.GinServer.Engine.Group("/api/v1/auth")
	{
		userService := &userservice.UserService{Fw: fw, Collection: "user"}
		authController := &authcontroller.AuthController{UserService: *userService}

		//Creates new user
		v1AuthRoutes.POST("/refreshToken/",

			//Valided need params and set the incoming data to the model
			//we then use it in controller
			middlewares.Check4ValidData(&authController.AuthRefreshTokenModel),
			authController.GenerateNewAccessTokenByRefreshToken)

	}
}
