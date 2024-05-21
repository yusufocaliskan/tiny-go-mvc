package authcontroller

import (
	"github.com/gin-gonic/gin"
	authmodel "github.com/gptverse/init/app/models/auth-model"
	usermodel "github.com/gptverse/init/app/models/user-model"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/framework/http/responser"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
	"github.com/gptverse/init/framework/translator"
)

type AuthController struct {
	// users being binded in IsValidate()
	UserService           userservice.UserService
	AuthRefreshTokenModel authmodel.AuthRefreshTokenModel
}

// Generating new accessToken using refreshToken
func (authCtrl *AuthController) GenerateNewAccessTokenByRefreshToken(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	data, _ := tinytoken.VerifyToken(authCtrl.AuthRefreshTokenModel.RefreshToken, authCtrl.UserService.Fw.Configs.AUTH_TOKEN_SECRET_KEY)

	//extract user data from bearer token
	userEmail := data["data"].(string)

	// Is user exists?
	isExists, user := authCtrl.UserService.CheckByEmailAddress(userEmail)

	// User Exists
	if !isExists {
		response.SetMessage(translator.GetMessage(ginCtx, "user_not_found")).BadWithAbort()
		return
	}

	//Generate new tokens.
	token := tinytoken.TinyToken{
		SecretKey: authCtrl.UserService.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}

	token.GenerateAccessTokens(user.Email)
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  user,
	}

	//return the resonse
	response.Payload(payload).Success()

}
