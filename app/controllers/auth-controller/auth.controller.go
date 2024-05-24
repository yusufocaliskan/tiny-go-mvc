package authcontroller

import (
	"log"

	"github.com/gin-gonic/gin"
	authmodel "github.com/gptverse/init/app/models/auth-model"
	usermodel "github.com/gptverse/init/app/models/user-model"
	authservice "github.com/gptverse/init/app/service/auth-service"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/framework/http/responser"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
	"github.com/gptverse/init/framework/translator"
)

type AuthController struct {
	UserService           userservice.UserService
	AuthService           authservice.AuthService
	AuthLoginModel        authmodel.AuthLoginModel
	AuthRefreshTokenModel authmodel.AuthRefreshTokenModel
}

//	@Tags			Auth
//	@Summary		Login
//	@Description	Sing-in With Access Token
//	@ID				access-token-login
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	tinytoken.TinyTokenSwaggerStruct
//	@Param			request			body		authmodel.AuthRefreshTokenModel	true	"query params"
//	@Param			Accept-Language	header		string							false	"Language preference"
//
//	@Router			/api/v1/auth/login [post]
func (authCtrl *AuthController) LoginWithAccessToken(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	user, isExists := authCtrl.UserService.GetUserByEmailAndPassword(authCtrl.AuthLoginModel.Email, authCtrl.AuthLoginModel.Password)

	if !isExists {
		log.Printf("Failed login attempt for email: %s", authCtrl.AuthLoginModel.Email)

		response.SetMessage(translator.GetMessage(ginCtx, "user_not_found")).BadWithAbort()
		return
	}
	token := tinytoken.TinyToken{
		SecretKey: authCtrl.AuthService.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}

	token.GenerateAccessTokens(authCtrl.AuthLoginModel.Email)
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  user.ToUserWithoutPassword(),
	}

	response.Payload(payload).Success()

}

//	@Tags			Auth
//	@Summary		Refresh Token
//	@Description	Generating new accessToken using refreshToken
//	@ID				refresh-token
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	tinytoken.TinyTokenSwaggerStruct
//	@Param			request			body		authmodel.AuthRefreshTokenModel	true	"query params"
//	@Param			Accept-Language	header		string							false	"Language preference"
//
//	@Router			/api/v1/auth/refreshToken [post]
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

	//return the resonse
	response.Payload(token.Data).Success()

}
