package tinytoken

import (
	"fmt"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TokenBlackList struct{}

// Add a token to the black list
func (bList *TokenBlackList) Add(ginCtx *gin.Context, token string, expireTime time.Time) {
	sesStore := sessions.Default(ginCtx)
	sesStore.Set(token, expireTime)
	sesStore.Save()
}

func (bList *TokenBlackList) IsBlackListed(ginCtx *gin.Context, token string) bool {
	sesStore := sessions.Default(ginCtx)

	expireTimeInterface := sesStore.Get(token)
	fmt.Println("IsBlackListed: OKKK--->-----", expireTimeInterface, token)
	if expireTimeInterface == nil {

		fmt.Println("OKKK--->-----", expireTimeInterface)
		return false
	}

	expireTime, ok := expireTimeInterface.(int64)
	fmt.Println("IsBlackListed : expireTime--->-----", expireTime)
	if !ok {

		fmt.Println("OKKK--->-----")
		return false
	}

	if time.Unix(expireTime, 0).Before(time.Now()) {
		// Clean the expired token
		sesStore.Delete(token)
		return false
	}

	fmt.Println("exists--->-----", expireTime)

	return true

	// if expireTimeInterface == nil {
	// 	return false
	// }

	// if expireTime.Before(time.Now()) {
	// 	sesStore.Delete(token)
	// 	return false

	// }

}
