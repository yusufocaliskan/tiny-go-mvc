package main

import (
	"encoding/gob"

	"github.com/gin-contrib/secure"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gptverse/init/app/middlewares"
	usermodel "github.com/gptverse/init/app/models/user-model"
	v1routes "github.com/gptverse/init/app/routes/v1"
	"github.com/gptverse/init/config"
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

//	@host	localhost:8080

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				"Bearer token for API authorization"
//	@description				Type "Bearer" followed by a space and JWT token.

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

	//Secure header
	secureConfigs := secure.New(secure.Config{
		AllowedHosts: config.AllowedHost,
		// SSLRedirect:  true,
		// SSLHost:               config.SSLHost,
		// SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval' https://cdnjs.cloudflare.com; style-src 'self' 'unsafe-inline' https://cdnjs.cloudflare.com; img-src 'self' data: https://cdnjs.cloudflare.com; font-src 'self' https://cdnjs.cloudflare.com;",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	})

	fw.GinServer.Engine.Static("/storage", "./storage")

	//Set the secure host
	fw.GinServer.Engine.Use(secureConfigs)

	//fetch user informations
	fw.GinServer.Engine.Use(middlewares.SetUserInformation2Session(&fw))

}

// Create a session store
func CreateSessionStore() {

	//Create session store using redis
	redisStore, _ := redis.NewStore(10, "tcp", confs.REDIS_DRIVER, "", []byte("secret"))
	fw.GinServer.Engine.Use(sessions.Sessions(confs.SESSION_KEY_NAME, redisStore))

	// Register the types with gob
	gob.Register(usermodel.UserModel{})
}

// Loads the v1 routes
func LoadV1Routes() {
	v1routes.SetAuthRoutes(&fw)
	v1routes.SetUserRoutes(&fw)
	v1routes.SetSwaggerRoute(&fw)
	v1routes.SetFileManagerRoutes(&fw)

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
