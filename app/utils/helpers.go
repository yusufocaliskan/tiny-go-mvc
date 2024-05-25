package utils

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	usermodel "github.com/gptverse/init/app/models/user-model"
)

// Check if array contains a value
func IsContains[T comparable](arr []T, x T) bool {
	for _, v := range arr {
		if v == x {
			return true
		}
	}
	return false
}

func GetCurrentUserInformations(ctx *gin.Context) *usermodel.UserModel {

	sesStore := sessions.Default(ctx)
	fetchCurrentUserInfo := sesStore.Get("CurrentUserInformations")

	currentUserInfo, _ := fetchCurrentUserInfo.(*usermodel.UserModel)
	return currentUserInfo
}
