package v1routes

import (
	authcontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/auth-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/app/middlewares"
	authservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/auth-service"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
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
			// middlewares.ValidateAndBind(&authController.AuthLoginModel),
			middlewares.AuthCheck(fw, authService),

			authController.Logout)

		//Refrsh token
		v1AuthRoutes.POST("/refreshToken/",
			// middlewares.RateLimeter(),
			middlewares.ValidateAndBind(&authController.AuthRefreshTokenModel),

			authController.GenerateNewAccessTokenByRefreshToken)

	}
}
