package usermodel

import (
	"time"

	tinytoken "github.com/gptverse/init/framework/tiny-token"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName string             `json:"fullname"`
	UserName string             `json:"username" validate:"required"`

	Email          string    `json:"email" validate:"required,email"`
	Password       string    `json:"password,omitempty" validate:"required" bson:"-"`
	HashedPassword string    `json:"hashed_password" bson:"hashed_password"`
	CreatedAt      time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time `bson:"updated_at" json:"updated_at"`
	Role           string    `json:"role" validate:"required,oneof=admin moderator user"`

	Ip        string             `json:"ip" bson:"ip"`
	CreatedBy primitive.ObjectID `json:"created_by" bson:"created_by"`
}

type UserWithoutPasswordModel struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName  string             `json:"fullname"`
	UserName  string             `json:"username" validate:"required"`
	Email     string             `json:"email" validate:"required,email"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Role      string             `json:"role" validate:"required,oneof=admin moderator user"`
}

// Remove password.
func (u *UserModel) ToUserWithoutPassword() UserModel {
	return UserModel{
		Id:        u.Id,
		FullName:  u.FullName,
		UserName:  u.UserName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Role:      u.Role,
	}
}

type UserUpdateModel struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	FullName  string    `json:"fullname" bson:"fullname"`
	UserName  string    `json:"username" validate:"required" bson:"username"`
	Email     string    `json:"email" validate:"required,email" bson:"email"`
	Role      string    `json:"role" validate:"required,oneof=admin moderator user"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UserModelResponse struct {
	Id       primitive.ObjectID `json:"id"`
	FullName string             `json:"fullname"`
	UserName string             `json:"username"`
	Email    string             `json:"email" `
}

type UserDeleteModel struct {
	Id primitive.ObjectID `json:"id" validate:"required"`
}

type UserWithIDFormIDModel struct {
	Id string `form:"id" validate:"required"`
}

type UserFilterModel struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type UserWithToken struct {
	Token tinytoken.TinyTokenData `json:"tokens"`
	User  UserModel               `json:"user"`
}

type UserSwaggerParams struct {
	Email    string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
	Name     string `json:"name" example:"johndoe"`
	FullName string `json:"fullname" example:"John Doe"`
	UserName string `json:"username" example:"johndoe"`
	Role     string `json:"role" validate:"required,oneof=admin moderator user"`
}
