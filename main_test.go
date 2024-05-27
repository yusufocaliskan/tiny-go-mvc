package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework/loader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Main Suite
type MainTestSuite struct {
	suite.Suite
}

// SetupTest runs before each test
func (suite *MainTestSuite) SetupTest() {
	// Load configurations
	var ldr = loader.Loader{}
	confs = ldr.LoadEnvironments()
	fw.Configs = &confs

	//SEt debug mode to false
	os.Setenv("GIN_MODE", "release")
	gin.SetMode(gin.ReleaseMode)

	InitialTheTinyGoMvc()
	fw.GinServer.Engine.GET("/test-url/",
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
		})

}

// Is Environments Loaded?
func (suite *MainTestSuite) TestShouldLoadedEnvironments() {
	assert.NotNil(suite.T(), confs, "Configurations should be loaded")
}

// Is connected to Database?
func (suite *MainTestSuite) TestShouldConnect2MongoDatabase() {

	fw.Database.Connect()
	assert.NotNil(suite.T(), fw.Database, "Should connect to database")
}

func (suite *MainTestSuite) TestCheckIfUserExists() {

	uService := &userservice.UserService{Fw: &fw, Collection: database.UserCollectionName}

	isExists, user := uService.CheckByEmailAddress("yusuf21@gmail.com")
	assert.True(suite.T(), isExists, "Connot find user")
	fmt.Printf("User: %+v\n", user)
}

// Test Gin Server
func (suite *MainTestSuite) TestGinServer() {
	// Create the JSON body
	body := map[string]string{
		"email":    "yusuf21@gmail.com",
		"password": "212131",
	}
	bodyJSON, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(bodyJSON))
	req.Host = "localhost"

	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	fw.GinServer.Engine.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code, "Expected 200 status code for /api/v1/auth/login endpoint")
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
