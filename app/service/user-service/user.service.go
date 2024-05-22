package userservice

import (
	"context"
	"errors"
	"fmt"

	usermodel "github.com/gptverse/init/app/models/user-model"
	"github.com/gptverse/init/framework"
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
func (uSrv UserService) CheckByEmailAddress(email string) (bool, *usermodel.UserModel) {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	var user usermodel.UserModel
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
func (uSrv UserService) GetUserById(id string) (*usermodel.UserWitoutPasswordModel, bool) {

	userId, _ := primitive.ObjectIDFromHex(id)
	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"_id": userId})

	fmt.Println(result)
	var user usermodel.UserWitoutPasswordModel

	result.Decode(&user)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return &user, false
		}
	}

	return &user, true
}

// Get user by email address
func (uSrv UserService) GetUserByEmailAddress(email string) (bool, *usermodel.UserModel) {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	fmt.Println(result)
	var user usermodel.UserModel

	result.Decode(&user)

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
