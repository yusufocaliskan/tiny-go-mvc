package v1routes

import (
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

func SetUserRoutes(fw *framework.Framework) {
	v1 := fw.GinServer.Engine.Group("/api/v1/user")
	{
		uService := &userservice.UserService{Fw: fw}
		uCtrl := &usercontroller.UserController{Service: *uService}

		v1.POST("/create/", uCtrl.CreateNewUser)
	}
}
