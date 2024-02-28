package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Engine  *gin.Engine
	Address string
}

// Creates a new Gin server
func (gs *GinServer) CreateServer(ports ...int) {

	fmt.Println("------------ {Establishing a Gin Server.} ------------")
	r := gin.Default()
	gs.Engine = r
	port := 8080

	if len(ports) > 0 {
		port = ports[0]
	}

	if len(ports) <= 0 {
		log.Fatal("Port is not setted, the Gin is going to use 8080 port")
	}

	gs.Address = fmt.Sprintf(":%d", port)
	fmt.Println("------------ {The Gin Server has been created.} ------------")

}

// Start the initialized server
func (gs *GinServer) Start() {
	gs.Engine.Run(gs.Address)
}
