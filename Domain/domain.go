package domain

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Database
	TaskCollection string = "tasks"
	UserCollection string = "users"
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
	Database = mongoClient.Database("task_manager")
}

type Task struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	DueDate time.Time `bson:"due_date" json:"due_date"`
	Status string `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Role string `bson:"role" json:"role"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"-"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type RegisterUserInput struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}