package main

import (
	"context"

	v1routes "github.com/yusufocaliskan/tiny-go-mvc/app/routes/v1"
	"github.com/yusufocaliskan/tiny-go-mvc/database"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	// initial dbConncction
	dbInstance := database.MongoDatabase{}
	dbInstance.Address = "mongodb://127.0.0.1:27017/"
	dbInstance.DBName = "tinyGoMvc"
	dbInstance.Connect()

	//Test
	ctx := context.Background()
	coll := dbInstance.Database.Collection("user")
	coll.InsertOne(ctx, bson.M{"mail": "test@mail.com"})

	//Create Gin Server
	RunGinServer()
}

// Runing the Gin Server
func RunGinServer() {

	//Create the server
	gsServer := &server.GinServer{}
	gsServer.CreateServer(4141)

	//Load routes
	LoadV1Routes(gsServer)

	//Start it
	gsServer.Start()
}

// Loads the v1 routes
func LoadV1Routes(gsServer *server.GinServer) {
	v1routes.SetUserRoutes(gsServer)
}
