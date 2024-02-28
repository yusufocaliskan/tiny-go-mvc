package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userservice "github.com/yusufocaliskan/tiny-go-mvc/app/service/user-service"
)

type UserController struct {
	Service userservice.UService
}

type User struct {
	FullName string
	LastName string
	Email    string
	Password string
}

// Get user by id
func (uCtrl *UserController) GetUserId(ginCtx *gin.Context) {

	// coll := uCtrl.Fw.Database.Instance.Collection("test")
	// coll.InsertOne(context.Background(), bson.M{"newInstanceeee": "TestooInstan"})
	uCtrl.Service.CreateNewUser()

	// fmt.Println("Testooooooo==?", coll)
	ginCtx.JSON(http.StatusOK, gin.H{"userId": uCtrl.Service.Fw.Database.DBName})
}
