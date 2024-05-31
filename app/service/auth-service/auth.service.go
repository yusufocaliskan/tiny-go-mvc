package authservice

import (
	"context"
	"errors"
	"time"

	authmodel "github.com/yusufocaliskan/tiny-go-mvc/app/models/auth-model"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// check auth.route
type AuthService struct {
	Collection string // user_auth
	Fw         *framework.Framework
}

func (Srv *AuthService) SaveToken(ctx context.Context, token *tinytoken.TinyTokenData, email string, status string) (bool, int64) {

	filter := bson.M{"email": email}

	updateData := bson.M{
		"$set": bson.M{
			"token":      token,
			"email":      email,
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

func (Srv *AuthService) SetTokenStatus(status string, email string) (bool, int64) {

	ctx := context.Background()
	filter := bson.M{"email": email}

	updateData := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	coll := Srv.Fw.Database.Instance.Collection(Srv.Collection)

	result, err := coll.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return false, 0
	}

	return true, result.ModifiedCount
}

// Check if user exists by given email address
func (Srv AuthService) GetToken(email string) (bool, *authmodel.AuthUserTokenModel) {

	ctx := context.Background()

	coll := Srv.Fw.Database.Instance.Collection(Srv.Collection)
	result := coll.FindOne(ctx, bson.M{"email": email})

	var user authmodel.AuthUserTokenModel
	result.Decode(&user)

	//User not found
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {

			return false, &user
		}
	}

	return true, &user
}
