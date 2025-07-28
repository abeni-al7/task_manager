package repositories

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
)

func ConnectToMongoDB() {
	mongoClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	db := mongoClient.Database("task_manager")
	TaskCollection = db.Collection("tasks")
	UserCollection = db.Collection("users")
}