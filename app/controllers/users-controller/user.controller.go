package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId int
}

// Get user by id
func (uCtrl *User) GetUserId(ginCtx *gin.Context) {
	ginCtx.JSON(http.StatusOK, gin.H{"userId": 212131})
}
