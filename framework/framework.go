package framework

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework/loader"
	"github.com/gptverse/init/framework/server"
)

type Framework struct {
	GinServer    *server.GinServer
	Database     *database.MongoDatabase
	Configs      *loader.Envs
	RedisStore   *redis.Store
	SessionStore *redis.Store
}
