package v1routes

import (
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetSwaggerRoute(fw *framework.Framework) {

	v1SwaggerRoutes := fw.GinServer.Engine.Group("/swagger")
	{

		v1SwaggerRoutes.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}

}
