package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	DbClient *mongo.Client
	Instance *mongo.Database
	DbUri    string
	DBName   string
}

// Make a Mongo db Connection
func (mongDb *MongoDatabase) Connect() {

	fmt.Println("------------ {Establising Database Connection} ------------ ")

	ctx := context.Background()

	//Database connection
	clientOptions := options.Client().ApplyURI(mongDb.DbUri)
	client, err := mongo.Connect(ctx, clientOptions)
	// defer client.Disconnect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	//check if the database accessable
	// isConnected := client.Ping(context.Background(), nil)
	// if isConnected != nil {
	// 	log.Fatal("Failed to Ping database", isConnected)
	// }

	db := client.Database(mongDb.DBName)

	//set the to the struct
	mongDb.DbClient = client
	mongDb.Instance = db

}
