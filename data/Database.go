package data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var MongoURI string = os.Getenv("MONGO_URI")
var Client *mongo.Client
var databaseName string = os.Getenv("MONGO_DB")


var Database *mongo.Database

func Configure() {
	log.Print("Configuring MongDB Connection")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(MongoURI))
	Database = Client.Database(databaseName)

	log.Println("Connection to MongoDB Established")
}