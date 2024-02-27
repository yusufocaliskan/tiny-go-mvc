package v1routes

import (
	usercontroller "github.com/yusufocaliskan/tiny-go-mvc/app/controllers/users-controller"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
)

func SetUserRoutes(ginServer *server.GinServer) {
	v1 := ginServer.Engine.Group("/v1/api/user")
	{
		uController := &usercontroller.User{}
		v1.GET("/getById/:Id", uController.GetUserId)
	}
}
