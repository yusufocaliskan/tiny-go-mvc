package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	tinyresponse "github.com/yusufocaliskan/tiny-go-mvc/framework/http/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	User    usermodel.UserModel
	Service userservice.UserService
}

var validate = validator.New()

// Get user by id
func (uCtrl *UserController) CreateNewUser(ginCtx *gin.Context) {

	Response := tinyresponse.Response{Ctx: ginCtx}

	//Set the data in comming data to the user model
	ginCtx.BindJSON(&uCtrl.User)

	//check if the validation is okay
	validationError := validate.Struct(&uCtrl.User)
	if validationError != nil {
		Response.Bad(validationError)
		return
	}

	//Genetate Id
	uCtrl.User.Id = primitive.NewObjectID()
	uCtrl.User.CreatedAt = time.Now()

	//Create new user
	uCtrl.Service.CreateNewUser(&uCtrl.User)

	Response.Success(uCtrl.User)
}
