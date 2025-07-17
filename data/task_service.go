package data

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/abeni-al7/task_manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var collection *mongo.Collection

func ConnectToMongoDB() {
	mongoClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
	collection = mongoClient.Database("task_manager").Collection("tasks")
}

func GetTasksService() ([]models.Task, error) {
	var tasks []models.Task

	cur, err := collection.Find(context.TODO(), bson.D{{}})
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

	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

func UpdateTaskService(id primitive.ObjectID, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{}}
	if updatedTask.Title != "" {
		update = append(update, bson.E{Key: "title", Value: updatedTask.Title})
	}
	if updatedTask.Description != "" {
		update = append(update, bson.E{Key: "description", Value: updatedTask.Description})
	}
	if !time.Time.IsZero(updatedTask.DueDate) {
		update = append(update, bson.E{Key: "due_date", Value: updatedTask.DueDate})
	}
	if updatedTask.Status != "" {
		update = append(update, bson.E{Key: "status", Value: updatedTask.Status})
	}
	update = append(update, bson.E{Key: "updated_at", Value: time.Now()})

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	
	err = collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func RemoveTaskService(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("task not found")
	}

	return nil
}

func AddTaskService(newTask models.Task) (models.Task, error) {
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	_, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return models.Task{}, errors.New("cannot insert task")
	}
	return newTask, nil
}