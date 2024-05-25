package filemanagermodel

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileModel struct {
	File *multipart.FileHeader `form:"file" validate:"required"`
}

type FileManagerUploadeFileResponse struct {
	Url string `json:"url" bson:"url"`
}

type FileDatabaseModel struct {
	Url      string             `json:"url" bson:"url"`
	UserId   primitive.ObjectID `json:"userId" bson:"userId"`
	CreateAt time.Time          `json:"created_at" bson:"created_at"`
	File     FileModel          `json:"file" bson:"file"`
}
