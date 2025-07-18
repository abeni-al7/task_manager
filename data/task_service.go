package data

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/abeni-al7/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
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

func GetTasksService() ([]models.Task, error) {
	var tasks []models.Task

	cur, err := TaskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []models.Task{}, errors.New("cannot retrieve tasks")
	}

	for cur.Next(context.TODO()) {
		var task models.Task
		err := cur.Decode(&task)
		if err != nil {
			return []models.Task{}, errors.New("cannot retrieve tasks")
		}
		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		return []models.Task{}, errors.New("cannot retrieve tasks")
	}

	cur.Close(context.TODO())

	return tasks, nil
}

func GetTaskService(id primitive.ObjectID) (models.Task, error) {
	var task models.Task

	filter := bson.D{{Key: "_id", Value: id}}

	err := TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

func UpdateTaskService(id primitive.ObjectID, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if updatedTask.Title != "" {
		fields = append(fields, bson.E{Key: "title", Value: updatedTask.Title})
	}
	if updatedTask.Description != "" {
		fields = append(fields, bson.E{Key: "description", Value: updatedTask.Description})
	}
	if !time.Time.IsZero(updatedTask.DueDate) {
		fields = append(fields, bson.E{Key: "due_date", Value: updatedTask.DueDate})
	}
	if updatedTask.Status != "" {
		fields = append(fields, bson.E{Key: "status", Value: updatedTask.Status})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err := TaskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, errors.New(err.Error())
	}
	
	err = TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func RemoveTaskService(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := TaskCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("task not found")
	}

	return nil
}

func AddTaskService(newTask models.Task) (models.Task, error) {
	newTask.ID = primitive.NewObjectID()
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	_, err := TaskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return models.Task{}, errors.New(err.Error())
	}
	return newTask, nil
}