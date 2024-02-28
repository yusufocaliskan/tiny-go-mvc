package userservice

import (
	"fmt"

	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

type UserService struct {
	Collection string
	User       *usermodel.UserModel
	Fw         *framework.Framework
}

func (uSrv *UserService) CreateNewUser() {

	fmt.Println("UService----", uSrv.Fw.Database.DBName)
}
