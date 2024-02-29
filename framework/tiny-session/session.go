package tinysession

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TinySession struct {
	Instance sessions.Session
}

func (ses *TinySession) New(ctx *gin.Context) sessions.Session {
	ses.Instance = sessions.Default(ctx)
	return ses.Instance
}

func (ses *TinySession) Get(key string) (bool, interface{}) {
	val := ses.Instance.Get(key)
	if val != nil {
		return true, val
	}

	return false, nil
}

func (ses *TinySession) Set(key string, val interface{}) {
	ses.Instance.Set(key, val)
}

func (ses *TinySession) Save(key string) {
	ses.Instance.Save()
}
