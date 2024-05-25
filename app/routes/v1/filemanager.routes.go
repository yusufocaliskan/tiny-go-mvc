package v1routes

import (
	filemanagercontroller "github.com/gptverse/init/app/controllers/filemanager-controller"
	"github.com/gptverse/init/app/middlewares"
	authservice "github.com/gptverse/init/app/service/auth-service"
	filemanagerservice "github.com/gptverse/init/app/service/filemanager-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework"
)

func SetFileManagerRoutes(fw *framework.Framework) {

	v1FileManagerRoutes := fw.GinServer.Engine.Group("/api/v1/file-manager")
	{
		//services
		fileService := &filemanagerservice.FileManagerService{Fw: fw, Collection: database.FileManagerCollectionName}
		authService := &authservice.AuthService{Fw: fw, Collection: database.AuthCollectionName}

		//userControler
		fileManagerController := &filemanagercontroller.FileManagerController{Service: *fileService}

		//Creates new user
		v1FileManagerRoutes.POST("/upload/",
			middlewares.ValidateAndBind(&fileManagerController.File),
			middlewares.AuthCheck(fw, authService),
			// middlewares.ForceOnlyRole("admin"),

			fileManagerController.Upload)

		//Get by Id
		// v1FileManagerRoutes.GET("/fetch",
		// 	middlewares.AuthCheck(fw, authService),
		// 	// middlewares.ValidateAndBind(&fileManagerController.UserWithIDFormIDModel),

		// 	fileManagerController.Fetch)

		//Get by Id
		v1FileManagerRoutes.GET("/fetch-all",
			middlewares.AuthCheck(fw, authService),
			// middlewares.ValidateAndBind(&fileManagerController.UserFilterModel),

			fileManagerController.FetchAll)

		// v1FileManagerRoutes.PUT("/update",
		// 	middlewares.AuthCheck(fw, authService),
		// 	// middlewares.ValidateAndBind(&fileManagerController.UserUpdateModel),

		// 	fileManagerController.Update)

		//Delete user
		//Only the one with {delete} permissions.
		v1FileManagerRoutes.DELETE("/delete/",
			middlewares.AuthCheck(fw, authService),
			// middlewares.ValidateAndBind(&fileManagerController.UserDeleteModel),

			fileManagerController.Delete)

	}
}
