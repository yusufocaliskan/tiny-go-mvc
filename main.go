package main

import (
	v1routes "github.com/yusufocaliskan/tiny-go-mvc/app/routes/v1"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
)

var fw = framework.Framework{}

func main() {

	//Initializing the framework
	//1. Start the Gin Framework, set it to the framework
	//2. Make database connection, set it to the framework
	InitialTheTinyGoMvc()
}

func InitialTheTinyGoMvc() {

	//1. Make database connection ad it to ginServer
	MongoDBConnection()

	//2. Create Gin Server
	RunGinServer()
}

// Runing the Gin Server
func RunGinServer() {

	ginServer := server.GinServer{}

	//Create the server
	ginServer.CreateServer(4141)

	//Set it the framwork
	fw.GinServer = &ginServer

	//Load routes
	LoadV1Routes()

	//Start it
	ginServer.Start()
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
	dbInstance.DbUri = "mongodb://127.0.0.1:27017/"
	dbInstance.DBName = "tinyGoMvc"
	dbInstance.Connect()

	fw.Database = &dbInstance

}
