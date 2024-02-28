package v1routes

import (
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

func SetUserRoutes(fw *framework.Framework) {
	v1 := fw.GinServer.Engine.Group("/v1/api/user")
	{
		uService := &userservice.UService{Fw: fw}
		uCtrl := &usercontroller.UserController{Service: *uService}

		v1.GET("/getById/:Id", uCtrl.GetUserId)
	}
}
