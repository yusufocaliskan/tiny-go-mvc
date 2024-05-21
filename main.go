package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gptverse/init/app/middlewares"
	v1routes "github.com/gptverse/init/app/routes/v1"
	"github.com/gptverse/init/database"
	_ "github.com/gptverse/init/docs"
	"github.com/gptverse/init/framework"
	"github.com/gptverse/init/framework/loader"
	"github.com/gptverse/init/framework/server"
)

//	@title			GPTVerse Admin Backend
//	@version		1.0
//	@description	To manage the whole gptv.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	@yusufocaliskan
//	@contact.email	yusufocaliskan@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host						localhost:4141
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

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

	// translation := translator.LoadErrorTextFile()

	// print("translation: ", translator.GetMessage(translation,"user_not_found"))

	//1. Make database connection ad it to ginServer
	MongoDBConnection()

	//2. Create Gin Server
	BuildGinServer()

	//3. Create session store using redis
	CreateSessionStore()

	//5. Middlewaress
	LoadMiddleWares()

	//5. Routes
	LoadV1Routes()

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

	//translation files
	fw.GinServer.Engine.Use(middlewares.LoadTranslationFile())

	//1. XSS attack protection
	fw.GinServer.Engine.Use(middlewares.AttackProtectionMiddleware())

	//Activating RateLimiiter.
	// if config.ActivateReteLimiter {

	// 	fmt.Println("-------- Rate Limitter is activated ----------")
	// 	fw.GinServer.Engine.Use(middlewares.RateLimeter())

	// }
}

// Create a session store
func CreateSessionStore() {

	//Create session store using redis
	redisStore, _ := redis.NewStore(10, "tcp", confs.REDIS_DRIVER, "", []byte("secret"))
	fw.GinServer.Engine.Use(sessions.Sessions(confs.SESSION_KEY_NAME, redisStore))

}

// Loads the v1 routes
func LoadV1Routes() {
	v1routes.SetAuthRoutes(&fw)
	v1routes.SetUserRoutes(&fw)
	v1routes.SetSwaggerRoute(&fw)

	//(..., call others here)
}

// Make db connection
func MongoDBConnection() {

	// initial dbConncction
	dbInstance := database.MongoDatabase{}

	dbInstance.DbUri = confs.DBUri
	dbInstance.DBName = confs.DBName
	dbInstance.DBPassword = confs.DB_PASSWORD
	dbInstance.DBUser = confs.DB_USER
	dbInstance.Connect()

	fw.Database = &dbInstance

}
