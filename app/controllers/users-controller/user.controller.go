package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
)

type UserController struct {
	User    usermodel.UserModel
	Service userservice.UserService
}

// Get user by id
func (uCtrl *UserController) GetUserId(ginCtx *gin.Context) {

	// coll := uCtrl.Fw.Database.Instance.Collection("test")
	// coll.InsertOne(context.Background(), bson.M{"newInstanceeee": "TestooInstan"})
	uCtrl.Service.CreateNewUser()

	// fmt.Println("Testooooooo==?", coll)
	ginCtx.JSON(http.StatusOK, gin.H{"userId": uCtrl.Service.Fw.Database.DBName})
}
