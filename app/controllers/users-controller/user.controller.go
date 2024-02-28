package usercontroller

import (
	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/Response"
)

type UserController struct {
	User    usermodel.UserModel
	Service userservice.UserService
}

// Get user by id
func (uCtrl *UserController) CreateNewUser(ginCtx *gin.Context) {

	Response := tinyresponse.Response{Ctx: ginCtx}
	err := ginCtx.ShouldBindJSON(&uCtrl.User)

	if err != nil {
		Response.Bad(err)
		return
	}

	//Create new user
	uCtrl.Service.CreateNewUser(&uCtrl.User)

	Response.Success(uCtrl.User)
}
