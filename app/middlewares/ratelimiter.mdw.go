package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gptverse/init/config"
	"github.com/gptverse/init/framework/http/request"
	responser "github.com/gptverse/init/framework/http/responser"
	tinysession "github.com/gptverse/init/framework/tiny-session"
	"github.com/gptverse/init/framework/translator"
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
		response := responser.Response{Ctx: ctx}

		sessionKey := fmt.Sprintf("limitterInfo-%s", clientIp)

		getSaveInformations := session.Get(sessionKey)

		fmt.Println("getSaveInformations------ ERRO", getSaveInformations)
		if getSaveInformations == nil {

			fmt.Println("getSaveInformations------ ERRO")

		}
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
					response.SetStatusCode(http.StatusRequestTimeout).SetMessage(translator.GetMessage(ctx, "reached_request_threshold")).BadWithAbort()
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
