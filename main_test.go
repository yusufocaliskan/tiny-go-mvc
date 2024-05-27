package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	authmodel "github.com/gptverse/init/app/models/auth-model"
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
	createSomeRoutes()
}

func createSomeRoutes() {

	fw.GinServer.Engine.POST("/set-session/",
		func(ctx *gin.Context) {
			var sessionModel authmodel.SessionModel
			sesStore := sessions.Default(ctx)

			if err := ctx.ShouldBindJSON(&sessionModel); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := sesStore.Save(); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
		})

	fw.GinServer.Engine.GET("/get-session/",
		func(ctx *gin.Context) {
			sesStore := sessions.Default(ctx)
			sesDataInterface := sesStore.Get("SesData")

			if sesDataInterface == nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "No session data found"})
				return
			}

			sessionModel, ok := sesDataInterface.(authmodel.SessionModel)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast session data"})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"sessionData": sessionModel})
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
func (suite *MainTestSuite) TestSessionStore() {
	// Step 1: Set session
	sessionData := authmodel.SessionModel{
		Email:    "test@mail.com",
		Password: "password123",
	}

	bodyJSON, _ := json.Marshal(sessionData)

	setSessionReq, _ := http.NewRequest("POST", "/set-session/", bytes.NewBuffer(bodyJSON))
	setSessionReq.Host = "localhost"
	setSessionReq.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	fw.GinServer.Engine.ServeHTTP(w1, setSessionReq)
	assert.Equal(suite.T(), http.StatusOK, w1.Code, "Expected 200 status code for /set-session/")

	// Step 2: Get session
	getSessionReq, _ := http.NewRequest("GET", "/get-session/", nil)

	getSessionReq.Host = "localhost"
	w2 := httptest.NewRecorder()
	fw.GinServer.Engine.ServeHTTP(w2, getSessionReq)
	assert.Equal(suite.T(), http.StatusOK, w2.Code, "Expected 200 status code for /get-session/")

	// Step 3: Parse and check the response
	var response struct {
		SessionData authmodel.SessionModel `json:"sessionData"`
	}
	err := json.Unmarshal(w2.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), sessionData.Email, response.SessionData.Email, "Expected session email to be 'test@mail.com'")
	assert.Equal(suite.T(), sessionData.Password, response.SessionData.Password, "Expected session password to be 'password123'")
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
