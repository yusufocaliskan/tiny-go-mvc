package usercontroller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	authservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/auth-service"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/request"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/responser"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/translator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	// users being binded in ValidateAndBind()
	User                     usermodel.UserModel
	UserDeleteModel          usermodel.UserDeleteModel
	UserWithIDFormIDModel    usermodel.UserWithIDFormIDModel
	UserWithoutPasswordModel usermodel.UserWithoutPasswordModel
	UserFilterModel          usermodel.UserFilterModel
	UserUpdateModel          usermodel.UserUpdateModel
	Service                  userservice.UserService
	AuthService              authservice.AuthService
}

// @Tags			Users
// @Summary		New user
// @Description	Creates new user
// @ID				create-user
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	usermodel.UserWithToken
// @Param			request			body		usermodel.UserSwaggerParams	true	"query params"
// @Param			Accept-Language	header		string						false	"Language preference"
//
// @Router			/api/v1/user/create [post]
func (uController *UserController) Create(ginCtx *gin.Context) {

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

	//Start e transaction
	ctx := context.TODO()

	dbSession, err := uController.AuthService.Fw.Database.Instance.Client().StartSession()
	if err != nil {
		log.Fatalf("Failed to start session %v", err)
	}
	//remove it when return
	defer dbSession.EndSession(ctx)

	//Generate tokens
	token := tinytoken.TinyToken{
		SecretKey: uController.Service.Fw.Configs.AUTH_TOKEN_SECRET_KEY,
	}

	err = mongo.WithSession(ctx, dbSession, func(sc mongo.SessionContext) error {
		err := dbSession.StartTransaction(options.Transaction())
		if err != nil {
			return fmt.Errorf("connot start the transaction")
		}

		token.GenerateAccessTokens(&uController.User.Email)

		//Genetate Id & Create new user
		uController.User.Id = primitive.NewObjectID()
		uController.User.Ip = request.GetLocalIP()
		uController.User.CreatedAt = time.Now()
		uController.User.CreatedBy = currentUserInfo.Id

		//hash the user's password
		hashedPassword, passwordHashingError := bcrypt.GenerateFromPassword([]byte(uController.User.Password), bcrypt.DefaultCost)
		uController.User.HashedPassword = string(hashedPassword)

		if passwordHashingError != nil {

			dbSession.AbortTransaction(sc)
			return fmt.Errorf("connot hash the password")
		}

		//Create new user
		_, isUserInserted := uController.Service.CreateNewUser(sc, uController.User)

		if !isUserInserted {
			dbSession.AbortTransaction(sc)
			return fmt.Errorf("connot create the user")
		}

		//save the token
		isTokenSaved, _ := uController.AuthService.SaveToken(sc, &token.Data, uController.User.Email, "active")

		if !isTokenSaved {
			dbSession.AbortTransaction(sc)
			return err
		}

		err = dbSession.CommitTransaction(sc)
		if err != nil {
			return fmt.Errorf("connot commit the transaction")
		}
		return nil

	})

	if err != nil {
		response.SetMessage(translator.GetMessage(ginCtx, "transaction_failed")).BadWithAbort()
		return
	}

	//Generate payload
	payload := usermodel.UserWithToken{
		Token: token.Data,
		User:  uController.User.ToUserWithoutPassword(),
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
// @Param			request			body		usermodel.UserUpdateSwaggerModel	true	"query params"
// @Param			Accept-Language	header		string								false	"Language preference"
//
// @Router			/api/v1/user/update [put]
func (uController *UserController) Update(ginCtx *gin.Context) {

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
		return
	}

	//Create new user
	isUpdated, user := uController.Service.UpdateUserInformations(newInformations, id)

	if !isUpdated {
		response.SetMessage(translator.GetMessage(ginCtx, "connot_update")).BadWithAbort()
		return
	}

	//return the resonse
	response.Payload(user).Success()

}

// @Tags			Users
// @Summary		Get User
// @Description	Get user details by id
// @ID				get-user-by-id
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	usermodel.UserWithoutPasswordModel
// @Param			id				query		string	true	"user id"	Format(ObjectID)
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/user/fetch [GET]
func (uController *UserController) Fetch(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	user, exists := uController.Service.GetUserById(uController.UserWithIDFormIDModel.Id)

	if !exists {
		response.SetMessage(translator.GetMessage(ginCtx, "user_not_found")).BadWithAbort()
		return
	}

	//return the resonse
	response.Payload(user.ToUserWithoutPassword()).Success()

}

// @Tags			Users
// @Summary		List All Records
// @Description	Get user details by id
// @ID				fetch-all-users
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	usermodel.UserWithoutPasswordModel
// @Param			page			query		string	true	"page number"	int
// @Param			limit			query		string	true	"limit number"	int
// @Param			Accept-Language	header		string	false	"Language preference"
//
// @Router			/api/v1/user/fetch-all [GET]
func (uController *UserController) FetchAll(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	filters := uController.UserFilterModel

	users, exists := uController.Service.FetchAll(filters)

	if !exists {
		response.SetMessage(translator.GetMessage(ginCtx, "user_not_found")).BadWithAbort()
		return
	}

	//return the resonse
	response.Payload(users).Success()

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
// @Router			/api/v1/user/delete [delete]
func (uController *UserController) Delete(ginCtx *gin.Context) {

	response := responser.Response{Ctx: ginCtx}

	isDeleted := uController.Service.DeleteUserById(&uController.UserDeleteModel)

	if !isDeleted {
		response.SetMessage(translator.GetMessage(ginCtx, "user_connot_delete")).BadWithAbort()
		return
	}

	response.SetMessage(translator.GetMessage(ginCtx, "user_deleted")).Success()

}
