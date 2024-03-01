package userservice

import (
	"context"
	"errors"
	"fmt"

	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// check user.route
type UserService struct {
	Collection string // user
	Fw         *framework.Framework
}

// Creeate a new user
func (uSrv *UserService) CreateNewUser(user *usermodel.UserModel) {

	ctx := context.Background()
	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)

	fmt.Println("Inserttting")
	coll.InsertOne(ctx, user)
	fmt.Println("Inserted")

}

// Check if user exists by given email address
func (uSrv UserService) CheckByEmailAddress(email string) bool {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {

			return false
		}
	}

	return true
}
