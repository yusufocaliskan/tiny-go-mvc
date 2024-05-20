package userservice

import (
	"context"
	"errors"
	"fmt"

	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	coll.InsertOne(ctx, user)

}

// Check if user exists by given email address
func (uSrv UserService) CheckByEmailAddress(email string) (bool, *usermodel.UserModelResponse) {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	var user usermodel.UserModelResponse
	result.Decode(&user)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {

			return false, &user
		}
	}

	return true, &user
}

// Check if user exists by given email address
func (uSrv UserService) GetUserById(id primitive.ObjectID) (bool, *usermodel.UserModelResponse) {

	ctx := context.Background()
	fmt.Println("USER ID ", id)

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"_id": id})

	var user usermodel.UserModelResponse
	result.Decode(&user)
	fmt.Println("user-->", result)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return false, &user
		}
	}

	return true, &user
}

// Delete user by id
func (uSrv UserService) DeleteUserById(data *usermodel.UserDeleteModel) bool {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result, _ := coll.DeleteOne(ctx, bson.M{"_id": data.Id})
	return result.DeletedCount > 0

}
