package authservice

import (
	"context"
	"errors"
	"time"

	usermodel "github.com/gptverse/init/app/models/user-model"
	"github.com/gptverse/init/framework"
	tinytoken "github.com/gptverse/init/framework/tiny-token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check auth.route
type AuthService struct {
	Collection string // user_auth
	Fw         *framework.Framework
}

func (Srv *AuthService) SaveToken(ctx context.Context, token *tinytoken.TinyTokenData, userId primitive.ObjectID, status string) (bool, int64) {

	filter := bson.M{"userId": userId}

	updateData := bson.M{
		"$set": bson.M{
			"token":      token,
			"userId":     userId,
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	coll := Srv.Fw.Database.Instance.Collection(Srv.Collection)

	opts := options.Update().SetUpsert(true)
	result, err := coll.UpdateOne(ctx, filter, updateData, opts)
	if err != nil {
		return false, 0
	}

	return true, result.ModifiedCount
}

// Check if user exists by given email address
func (Srv AuthService) GetToken(email string) (bool, *usermodel.UserModel) {

	ctx := context.Background()

	coll := Srv.Fw.Database.Instance.Collection(Srv.Collection)
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
