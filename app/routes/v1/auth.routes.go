package v1routes

import (
	authcontroller "github.com/gptverse/init/app/controllers/auth-controller"
	"github.com/gptverse/init/app/middlewares"
	authservice "github.com/gptverse/init/app/service/auth-service"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/framework"
)

func SetAuthRoutes(fw *framework.Framework) {
	v1AuthRoutes := fw.GinServer.Engine.Group("/api/v1/auth")
	{
		userService := &userservice.UserService{Fw: fw, Collection: "user"}
		authService := &authservice.AuthService{Fw: fw, Collection: "auth"}
		authController := &authcontroller.AuthController{UserService: *userService, AuthService: *authService}

		//Creates new user
		v1AuthRoutes.POST("/refreshToken/",

			middlewares.ValidateAndBind(&authController.AuthRefreshTokenModel),
			authController.GenerateNewAccessTokenByRefreshToken)

	}
}
