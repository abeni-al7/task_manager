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
	TaskCollection = mongoClient.Database("task_manager").Collection("tasks")
	UserCollection = mongoClient.Database("task_manager").Collection("users")
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

type TaskUsecase interface {
	Create(task *Task) (Task, error)
	FetchAll() ([]Task, error)
	Fetch(id primitive.ObjectID) (Task, error)
	Update(id primitive.ObjectID, updatedTask Task) (Task, error)
	Remove(id primitive.ObjectID) error
}

type UserUsecase interface {
	Register(user *User) (User, error)
	Login(username string, password string) (string, error)
	Promote(id primitive.ObjectID) (User, error)
	FetchAll() ([]User, error)
	Fetch(id primitive.ObjectID) (User, error)
	Update(id primitive.ObjectID, updatedUser User) (User, error)
	ChangePassword(id primitive.ObjectID, prevPassword string, newPassword string) error
	Remove(id primitive.ObjectID) error
}