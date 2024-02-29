package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	// users being binded in IsValidate()
	User    usermodel.UserModel
	Service userservice.UserService
}

func (uCtrl *UserController) CreateNewUser(ginCtx *gin.Context) {

	Response := tinyresponse.Response{Ctx: ginCtx}

	//Genetate Id & Create new user
	uCtrl.User.Id = primitive.NewObjectID()
	uCtrl.User.CreatedAt = time.Now()

	uCtrl.Service.CreateNewUser(&uCtrl.User)
	Response.Success(uCtrl.User)

}
