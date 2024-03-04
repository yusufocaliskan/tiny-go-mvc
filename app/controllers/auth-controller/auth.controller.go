package authcontroller

import (
	"strings"

	"github.com/gin-gonic/gin"
	errormessages "github.com/yusufocaliskan/tiny-go-mvc/app/constants/error-messages"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
)

type AuthController struct {
	// users being binded in IsValidate()
	UserService userservice.UserService
}

// Generating new accessToken using refreshToken
func (authCtrl *AuthController) GenerateNewAccessTokenByRefreshToken(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	authHeader := ginCtx.GetHeader("Authorization")
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
	data, _ := tinytoken.VerifyToken(bearerToken, authCtrl.UserService.Fw.Configs.AUTH_TOKEN_SECRET_KEY)

	//extract user data from bearer token
	userEmail := data["data"].(string)

	// Is user exists?
	isExists, user := authCtrl.UserService.CheckByEmailAddress(userEmail)

	// User Exists
	if !isExists {
		response.SetError(errormessages.UserDoesntExists).BadWithAbort()
		return
	}

	//Generate new tokens.
	UserAccessTokens := tinytoken.TinyToken{
		SecretKey: authCtrl.UserService.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}
	UserAccessTokens.AccessTokenGenerator(user.Email)
	UserAccessTokens.RefreshTokenGenerator(user.Email)

	//return the resonse
	response.Payload(UserAccessTokens).Success()

}
