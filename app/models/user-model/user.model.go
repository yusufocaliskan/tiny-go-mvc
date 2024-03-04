package usermodel

import (
	"time"

	tinytoken "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	FullName  string              `json:"fullname"`
	UserName  string              `json:"username" validate:"required"`
	Email     string              `json:"email" validate:"required,email"`
	Password  string              `json:"password" validate:"required"`
	CreatedAt time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time           `bson:"updated_at" json:"updated_at"`
	Token     tinytoken.TinyToken `json:"tokens"`
}
