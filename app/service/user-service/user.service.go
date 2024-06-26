package userservice

import (
	"context"
	"errors"
	"fmt"
	"time"

	usermodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/user-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// check user.route
type UserService struct {
	Collection string // user
	Fw         *framework.Framework
}

// Creeate a new user
func (uSrv *UserService) CreateNewUser(ctx context.Context, user usermodel.UserModel) (interface{}, bool) {

	fmt.Println("USer serviceL-->", user.CreatedBy)
	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		return 0, false
	}

	return result.InsertedID, true

}

// Creeate a new user
func (uSrv *UserService) UpdateUserInformations(user *usermodel.UserUpdateModel, userId string) (bool, int64) {

	ctx := context.TODO()
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": id}

	updateData := bson.M{
		"$set": bson.M{
			"email":      user.Email,
			"fullname":   user.FullName,
			"username":   user.UserName,
			"role":       user.Role,
			"updated_at": time.Now(),
		},
	}

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result, err := coll.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return false, 0
	}

	return true, result.ModifiedCount

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
func (uSrv UserService) GetUserById(id string) (*usermodel.UserModel, bool) {

	userId, _ := primitive.ObjectIDFromHex(id)
	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"_id": userId})

	var user usermodel.UserModel

	result.Decode(&user)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return &user, false
		}
	}

	return &user, true
}

// Check if user exists by given email address and password
func (uSrv UserService) GetUserByEmailAndPassword(email string, password string) (*usermodel.UserModel, bool) {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	var user usermodel.UserModel

	result.Decode(&user)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return &user, false
		}
	}

	//Compare the hashd password
	isPasswordCorret := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if isPasswordCorret != nil {
		return nil, false
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

// Fetch all records
func (uSrv UserService) FetchAll(filters usermodel.UserFilterModel) ([]usermodel.UserWithoutPasswordModel, bool) {

	ctx := context.Background()
	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)

	//	Pagination
	if filters.Page <= 0 {
		filters.Page = 1
	}
	if filters.Limit <= 0 {
		filters.Limit = 15
	}

	skip := filters.Page*filters.Limit - filters.Limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(filters.Limit))

	//find all
	cursor, err := coll.Find(ctx, bson.D{}, findOptions)

	if err != nil {
		return nil, false
	}

	defer cursor.Close(ctx)

	var users []usermodel.UserWithoutPasswordModel
	if err := cursor.All(ctx, &users); err != nil {
		return nil, false
	}

	return users, true

}

// Delete user by id
func (uSrv UserService) DeleteUserById(data *usermodel.UserDeleteModel) bool {

	ctx := context.Background()

	coll := uSrv.Fw.Database.Instance.Collection(uSrv.Collection)
	result, _ := coll.DeleteOne(ctx, bson.M{"_id": data.Id})
	return result.DeletedCount > 0

}
