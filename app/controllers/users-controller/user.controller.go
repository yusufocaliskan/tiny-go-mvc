package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	errormessages "github.com/yusufocaliskan/tiny-go-mvc/app/constants/error-messages"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinyerror "github.com/yusufocaliskan/tiny-go-mvc/framework/http/tiny-error"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	// users being binded in IsValidate()
	User    usermodel.UserModel
	Service userservice.UserService
}

func (uController *UserController) CreateNewUserByEmailAdress(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	//Is user exists?
	isExists := uController.Service.CheckByEmailAddress(uController.User.Email)

	if isExists {

		response.Error = tinyerror.New(errormessages.UserExists)
		response.BadWithAbort()
		return
	}

	//Genetate Id & Create new user
	uController.User.Id = primitive.NewObjectID()
	uController.User.CreatedAt = time.Now()

	uController.Service.CreateNewUser(&uController.User)
	response.Payload(uController.User).Success()

}
