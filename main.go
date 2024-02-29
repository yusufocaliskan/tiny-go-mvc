package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/yusufocaliskan/tiny-go-mvc/app/middlewares"
	v1routes "github.com/yusufocaliskan/tiny-go-mvc/app/routes/v1"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/loader"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
)

var fw = framework.Framework{}

// Loadd the configuratins
var ldr = loader.Loader{}
var confs = ldr.LoadEnvironmetns()

func main() {

	//Initializing the framework
	InitialTheTinyGoMvc()
}

// Workflow of the app..
func InitialTheTinyGoMvc() {

	//1. Make database connection ad it to ginServer
	MongoDBConnection()

	//2. Create Gin Server
	BuildGinServer()

	//3. Routes
	LoadV1Routes()

	//4. Create session store using redis
	CreateSessionStore()

	//5. Middlewaress
	LoadMiddleWares()

	//Start it 🚀
	fw.GinServer.Start()
}

// Runing the Gin Server
func BuildGinServer() {

	ginServer := server.GinServer{}

	//Create the server
	ginServer.CreateServer(confs.GIN_SERVER_PORT)

	//Set it the framwork
	fw.GinServer = &ginServer

	//set the configurationss
	fw.Configs = &confs

}

// Set your general middleware
func LoadMiddleWares() {
	fw.GinServer.Engine.Use(middlewares.RateLimeter())
}

// Create a session store
func CreateSessionStore() {

	//Create session store using redis
	redisStore, _ := redis.NewStore(10, "tcp", confs.REDIS_DRIVER, "", []byte("secret"))
	fw.GinServer.Engine.Use(sessions.Sessions(confs.SESSION_KEY_NAME, redisStore))

}

// Loads the v1 routes
func LoadV1Routes() {
	v1routes.SetUserRoutes(&fw)

	//(..., call others here)
}

// Make db connection
func MongoDBConnection() {

	// initial dbConncction
	dbInstance := database.MongoDatabase{}

	//TODO: Get the frome .env
	dbInstance.DbUri = confs.DBUri
	dbInstance.DBName = confs.DBName
	dbInstance.Connect()

	fw.Database = &dbInstance

}

func CreateRedisStore() {
}
