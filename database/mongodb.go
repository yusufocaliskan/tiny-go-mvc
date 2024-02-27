package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Client   *mongo.Client
	Database *mongo.Database
	Address  string
	DBName   string
}

// Make a Mongo db Connection
func (mongDb *MongoDatabase) Connect() {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongDb.Address))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(mongDb.DBName)

	//set the to the struct
	mongDb.Client = client
	mongDb.Database = db

}
