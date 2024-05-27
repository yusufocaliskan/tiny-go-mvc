package main

import (
	"fmt"
	"testing"

	userservice "github.com/gptverse/init/app/service/user-service"
	"github.com/gptverse/init/database"
	"github.com/gptverse/init/framework/loader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock Database
type MockDatabase struct {
	mock.Mock
	Instance *mongo.Database
}

// Main Suite
type MainTestSuite struct {
	suite.Suite
}

var (
	mockDB *MockDatabase
)

func (m *MockDatabase) Connect() {
	m.Called()
}

// SetupTest runs before each test
func (suite *MainTestSuite) SetupTest() {
	// Load configurations
	var ldr = loader.Loader{}
	confs = ldr.LoadEnvironments()
	fw.Configs = &confs

	// Mock the database connection
	mockDB = new(MockDatabase)
	mockDB.On("Connect").Return(nil)

	dbInstance := &database.MongoDatabase{
		DbUri:      confs.DBUri,
		DBName:     confs.DBName,
		DBUser:     confs.DB_USER,
		DBPassword: confs.DB_PASSWORD,
	}

	// Assuming database.Connect sets the instance field
	dbInstance.Connect()
	fw.Database = dbInstance
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

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
