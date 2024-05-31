package framework

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/yusufocaliskan/tiny-go-mvc/config"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
)

type Framework struct {
	GinServer    *server.GinServer
	Database     *database.MongoDatabase
	Configs      *config.Envs
	RedisStore   *redis.Store
	SessionStore *redis.Store
}
