package userservice

import (
	"context"

	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/model/user-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
)

type UserService struct {
	Collection string // user
	Fw         *framework.Framework
}

// Creeate a new user
func (uSrv *UserService) CreateNewUser(user *usermodel.UserModel) {
	ctx := context.Background()
	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	coll.InsertOne(ctx, user)

}
