package usercontroller

import (
	"time"

	"github.com/gin-gonic/gin"
	errormessages "github.com/yusufocaliskan/tiny-go-mvc/app/constants/error-messages"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	// users being binded in IsValidate()
	User    usermodel.UserModel
	Service userservice.UserService
}

func (uController *UserController) CreateNewUserByEmailAdress(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}
	UserAccessTokens := tinytoken.TinyToken{
		SecretKey: uController.Service.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}
	UserAccessTokens.AccessTokenGenerator(&uController.User)
	UserAccessTokens.RefreshTokenGenerator(&uController.User)

	//Is user exists?
	isExists := uController.Service.CheckByEmailAddress(uController.User.Email)

	// User Exists
	if isExists {
		response.SetError(errormessages.UserExists).BadWithAbort()
		return
	}

	//Genetate Id & Create new user
	uController.User.Id = primitive.NewObjectID()
	uController.User.CreatedAt = time.Now()
	uController.User.Token = UserAccessTokens

	//Create new user
	uController.Service.CreateNewUser(&uController.User)

	//return the resonse
	response.Payload(uController.User).Success()

}
