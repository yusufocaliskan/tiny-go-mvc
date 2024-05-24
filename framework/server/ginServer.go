package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	Engine  *gin.Engine
	Address string
	Port    int
}

// Creates a new Gin server
func (gs *GinServer) CreateServer(port int) {

	fmt.Println("------------ {Establishing a Gin Server.} ------------")
	r := gin.Default()

	gs.Engine = r
	gs.Port = port

	gs.Address = fmt.Sprintf(":%d", port)
	fmt.Println("------------ {The Gin Server has been created.} ------------")

}

// Start the initialized server
func (gs *GinServer) Start() {
	gs.Engine.Run(gs.Address)
}
