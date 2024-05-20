package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	// users being binded in Check4ValidData()
	User            usermodel.UserModel
	UserDeleteModel usermodel.UserDeleteModel
	Service         userservice.UserService
}

func (uController *UserController) CreateNewUserByEmailAdress(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}
	//Is user exists?
	isExists, _ := uController.Service.CheckByEmailAddress(uController.User.Email)

	// User Exists
	if isExists {

		useExistsError := translator.GetMessage(ginCtx, "user_exists")

		response.SetMessage(useExistsError).BadWithAbort()
		return
	}

	//Generate tokens
	token := tinytoken.TinyToken{
		SecretKey: uController.Service.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}
	token.GenerateAccessTokens(&uController.User.Email)

	//Genetate Id & Create new user
	uController.User.Id = primitive.NewObjectID()
	uController.User.CreatedAt = time.Now()

	//Create new user
	uController.Service.CreateNewUser(&uController.User)

	//Generate payload
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  &uController.User,
	}

	//return the resonse
	response.Payload(payload).Success()

}

// Deletes a user
func (uController *UserController) DeleteUserById(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	isDeleted := uController.Service.DeleteUserById(&uController.UserDeleteModel)

	if !isDeleted {
		response.SetMessage(translator.GetMessage(ginCtx, "connot_delete_user")).BadWithAbort()
		return
	}

	response.SetMessage(translator.GetMessage(ginCtx, "user_deleted")).Success()

}
