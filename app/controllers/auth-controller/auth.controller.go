package authcontroller

import (
	"github.com/gin-gonic/gin"
	textholder "github.com/yusufocaliskan/tiny-go-mvc/app/constants/text-holder/eng"
	authmodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/auth-model"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
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
		response.SetError(textholder.UserDoesntExists).BadWithAbort()
		return
	}

	//Generate new tokens.
	token := tinytoken.TinyToken{
		SecretKey: authCtrl.UserService.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}

	token.GenerateAccessTokens(user.Email)
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  *user,
	}

	//return the resonse
	response.Payload(payload).Success()

}
