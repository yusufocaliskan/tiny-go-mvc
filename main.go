package main

import (
	v1routes "github.com/yusufocaliskan/tiny-go-mvc/app/routes/v1"
	"github.com/yusufocaliskan/tiny-go-mvc/framework/server"
)

func main() {

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
