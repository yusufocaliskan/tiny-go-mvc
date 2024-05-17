package middlewares

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	// textholder "github.com/yusufocaliskan/tiny-go-mvc/app/constants/text-holder/eng"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/http/request"
	tinysession "github.com/yusufocaliskan/tiny-go-mvc/framework/tiny-session"
)

type Limiter struct {
	LastRequestTime time.Time
	Endpoint        url.URL
	Token           int
	ClientIp        string
}

func RateLimeter() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("-------- Rate Limits works ----------")

		tinySession := &tinysession.TinySession{}
		session := tinySession.New(ctx)
		clientIp := request.GetLocalIP()
		// response := responser.Response{Ctx: ctx}

		sessionKey := fmt.Sprintf("limitterInfo-%s", clientIp)

		getSaveInformations := session.Get(sessionKey)
		var limiterInfo Limiter
		var limiterInfoAsJson []byte
		var remainedTime time.Duration
		var elapsedTime time.Duration
		var rateLimit time.Duration

		//we have saved info?
		if getSaveInformations != nil {

			if err := json.Unmarshal([]byte(getSaveInformations.(string)), &limiterInfo); err != nil {
				// Handle error
				fmt.Println("Error decoding limiterInfo JSON:", err)
			}

			elapsedTime = time.Since(limiterInfo.LastRequestTime)
			rateLimit = config.RateLimterTime * time.Second
			remainedTime = rateLimit - elapsedTime

			//we still have token.
			if limiterInfo.Token > 0 {
				limiterInfo.Token--
			}

			//We have reached the threshold
			if limiterInfo.Token < config.RateLimiterToken {

				//is ther remainedTime?
				if remainedTime <= 0 && limiterInfo.Token <= 0 {
					limiterInfo = Limiter{
						LastRequestTime: time.Now(),
						Token:           config.RateLimiterToken,
						ClientIp:        clientIp,
						Endpoint:        *ctx.Request.URL,
					}
				}

				if limiterInfo.Token == 0 {
					// message := fmt.Sprintf(textholder.RateLimiterThresholMessage, remainedTime.Minutes())
					// response.SetStatusCode(http.StatusRequestTimeout).SetError(message).BadWithAbort()
				}

			}

			limiterInfoAsJson, _ := json.Marshal(limiterInfo)

			//if there is still token, then decreas some..
			session.Set(sessionKey, string(limiterInfoAsJson))
			session.Save()

		}

		//if we dont'have saved information
		if getSaveInformations == nil {
			//we don't have saved data, crate new one
			limiterInfo = Limiter{
				LastRequestTime: time.Now(),
				Token:           config.RateLimiterToken,
				ClientIp:        clientIp,
				Endpoint:        *ctx.Request.URL,
			}

			limiterInfoAsJson, _ = json.Marshal(limiterInfo)
			session.Set(sessionKey, string(limiterInfoAsJson))
			session.Save()
		}

		// Proceed with the request
		ctx.Next()
	}
}
