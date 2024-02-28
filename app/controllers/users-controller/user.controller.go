package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	form "github.com/yusufocaliskan/tiny-go-mvc/framework/form/validate"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	User    usermodel.UserModel
	Service userservice.UserService
}

// Get user by id
func (uCtrl *UserController) CreateNewUser(ginCtx *gin.Context) {

	Response := tinyresponse.Response{Ctx: ginCtx}
	validate := form.FormValidator{}

	//Set the data in comming data to the user model
	ginCtx.BindJSON(&uCtrl.User)

	//check if the validation is okay
	validationErrors := validate.Check(&uCtrl.User)
	if validationErrors != nil {
		Response.Bad(validationErrors)
		return
	}

	//Genetate Id & Create new user
	uCtrl.User.Id = primitive.NewObjectID()
	uCtrl.User.CreatedAt = time.Now()

	uCtrl.Service.CreateNewUser(&uCtrl.User)

	Response.Success(uCtrl.User)
}
