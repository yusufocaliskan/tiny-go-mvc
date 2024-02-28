package userservice

import (
	"fmt"

	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

type User struct {
	FullName string
	LastName string
	Email    string
	Password string
}

type UserService struct {
	Collection string
	User       User
	Fw         *framework.Framework
}

func (uSrv *UserService) CreateNewUser() {

	fmt.Println("UService----", uSrv.Fw.Database.DBName)
}
