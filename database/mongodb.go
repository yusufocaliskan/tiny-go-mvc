package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	DbClient   *mongo.Client
	Instance   *mongo.Database
	DbUri      string
	DBName     string
	DBUser     string
	DBPassword string
}

// Make a Mongo db Connection
func (mongDb *MongoDatabase) Connect() {

	fmt.Println("------------ {Establising Database Connection} ------------ ")

	fmt.Println("Database mongDb.DbUri:--> ", mongDb.DbUri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Set auth
	// authCredential := options.Credential{
	// 	Username: mongDb.DBUser,
	// 	Password: mongDb.DBPassword,
	// }

	//Database connection
	clientOptions := options.Client().ApplyURI(mongDb.DbUri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := client.Database(mongDb.DBName)

	//set the to the struct
	mongDb.DbClient = client
	mongDb.Instance = db

}
