package usercontroller

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	usermodel "github.com/gptverse/init/app/models/user-model"
	authservice "github.com/gptverse/init/app/service/auth-service"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/framework/http/request"
	"github.com/gptverse/init/framework/http/responser"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
	"github.com/gptverse/init/framework/translator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	// users being binded in ValidateAndBind()
	User                    usermodel.UserModel
	UserDeleteModel         usermodel.UserDeleteModel
	UserWithIDFormIDModel   usermodel.UserWithIDFormIDModel
	UserWitoutPasswordModel usermodel.UserWitoutPasswordModel
	UserUpdateModel         usermodel.UserUpdateModel
	Service                 userservice.UserService
	AuthService             authservice.AuthService
}

// @Tags			Users
// @Summary		New user
// @Description	Creates new user
// @ID				create-user
// @Accept			json
// @Produce		json
// @Success		200				{object}	usermodel.UserSwaggerParams
// @Param			id				query		string	true	"query params"
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/user/createByEmail [post]
func (uController *UserController) CreateNewUserByEmailAdress(ginCtx *gin.Context) {

	sesStore := sessions.Default(ginCtx)
	response := responser.Response{Ctx: ginCtx}
	//Is user exists?
	isExists, _ := uController.Service.CheckByEmailAddress(uController.User.Email)
	fetchCurrentUserInfo := sesStore.Get("CurrentUserInformations")

	currentUserInfo, _ := fetchCurrentUserInfo.(*usermodel.UserModel)

	if isExists {

		response.SetMessage(translator.GetMessage(ginCtx, "user_exists")).BadWithAbort()
		return
	}

	//No current user informations added to the Context
	if fetchCurrentUserInfo == nil {
		response.SetMessage(translator.GetMessage(ginCtx, "unknow_errors")).BadWithAbort()

		return
	}

	//Generate tokens
	token := tinytoken.TinyToken{
		SecretKey: uController.Service.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}
	token.GenerateAccessTokens(&uController.User.Email)

	//Genetate Id & Create new user
	uController.User.Id = primitive.NewObjectID()
	uController.User.Ip = request.GetLocalIP()
	uController.User.CreatedAt = time.Now()
	uController.User.CreatedBy = currentUserInfo.Id

	//Create new user
	uController.Service.CreateNewUser(&uController.User)

	//save the token
	uController.AuthService.SaveToken(&token.Data, uController.User.Id, "active")

	//Generate payload
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  &uController.User,
	}

	//return the resonse
	response.Payload(payload).Success()

}

// @Tags			Users
// @Summary		Update User
// @Description	Updates user informations by giving the Id
// @ID				update-user-informations
// @Accept			json
// @Produce		json
// @Success		200				{object}	usermodel.UserSwaggerParams
// @Param			request			body		usermodel.UserUpdateModel	true	"query params"
// @Param			Accept-Language	header		string						false	"Language preference"
//
// @Router			/api/v1/user/updateUserInformationsById [put]
func (uController *UserController) UpdateUserInformationsById(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}
	newInformations := &uController.UserUpdateModel

	id := newInformations.Id
	newInformations.UpdatedAt = time.Now()

	//Is user exists?
	_, isExists := uController.Service.GetUserById(id)

	// User Exists
	if !isExists {
		useExistsError := translator.GetMessage(ginCtx, "user_not_found")
		response.SetMessage(useExistsError).BadWithAbort()
	}

	//Create new user
	isUpdated, _ := uController.Service.UpdateUserInformations(newInformations, id)

	if !isUpdated {
		response.SetMessage(translator.GetMessage(ginCtx, "connot_update")).BadWithAbort()
		return
	}

	//return the resonse
	response.SetMessage(translator.GetMessage(ginCtx, "success_message")).Success()

}

// @Tags			Users
// @Summary		Get User
// @Description	Get user details by id
// @ID				get-user-by-id
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	usermodel.UserWitoutPasswordModel
// @Param			id				query		string	true	"user id"	Format(ObjectID)
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/user/getUserById [GET]
func (uController *UserController) GetUserById(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	user, exists := uController.Service.GetUserById(uController.UserWithIDFormIDModel.Id)

	if !exists {
		response.SetMessage(translator.GetMessage(ginCtx, "user_not_found")).BadWithAbort()
		return
	}

	//return the resonse
	response.SetMessage(translator.GetMessage(ginCtx, "success_message")).Payload(user).Success()

}

// @Tags			Users
// @Summary		Delete user
// @Description	Deletes a user by given user id
// @ID				Delete-User
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	translator.TranslationSwaggerResponse
// @Param			request			body		usermodel.UserDeleteModel	true	"query params"
// @Param			Accept-Language	header		string						false	"Language preference"
//
// @Router			/api/v1/user/deleteById [delete]
func (uController *UserController) DeleteUserById(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	isDeleted := uController.Service.DeleteUserById(&uController.UserDeleteModel)

	if !isDeleted {
		response.SetMessage(translator.GetMessage(ginCtx, "connot_delete_user")).BadWithAbort()
		return
	}

	response.SetMessage(translator.GetMessage(ginCtx, "user_deleted")).Success()

}
