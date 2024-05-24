package v1routes

import (
	authcontroller "github.com/gptverse/init/app/controllers/auth-controller"
	"github.com/gptverse/init/app/middlewares"
	authservice "github.com/gptverse/init/app/service/auth-service"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework"
)

func SetAuthRoutes(fw *framework.Framework) {
	v1AuthRoutes := fw.GinServer.Engine.Group("/api/v1/auth")
	{
		userService := &userservice.UserService{Fw: fw, Collection: database.UserCollectionName}

		authService := &authservice.AuthService{Fw: fw, Collection: database.AuthCollectionName}

		authController := &authcontroller.AuthController{UserService: *userService, AuthService: *authService}

		v1AuthRoutes.POST("/login",
			// middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&authController.AuthLoginModel),

			authController.LoginWithAccessToken)

		v1AuthRoutes.POST("/logout",
			// middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&authController.AuthLoginModel),
			middlewares.AuthCheck(fw),

			authController.LoginWithAccessToken)

		//Refrsh token
		v1AuthRoutes.POST("/refreshToken/",
			// middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&authController.AuthRefreshTokenModel),

			authController.GenerateNewAccessTokenByRefreshToken)

	}
}
