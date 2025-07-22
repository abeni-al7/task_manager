package data

import (
	"context"
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasksService() ([]domain.Task, error) {
	var tasks []domain.Task

	cur, err := domain.TaskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []domain.Task{}, errors.New("cannot retrieve tasks")
	}

	for cur.Next(context.TODO()) {
		var task domain.Task
		err := cur.Decode(&task)
		if err != nil {
			return []domain.Task{}, errors.New("cannot retrieve tasks")
		}
		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		return []domain.Task{}, errors.New("cannot retrieve tasks")
	}

	cur.Close(context.TODO())

	return tasks, nil
}

func GetTaskService(id primitive.ObjectID) (domain.Task, error) {
	var task domain.Task

	filter := bson.D{{Key: "_id", Value: id}}

	err := domain.TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}

	return task, nil
}

func UpdateTaskService(id primitive.ObjectID, updatedTask domain.Task) (domain.Task, error) {
	var task domain.Task
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

	_, err := domain.TaskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.Task{}, errors.New(err.Error())
	}
	
	err = domain.TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}
	return task, nil
}

func RemoveTaskService(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := domain.TaskCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("task not found")
	}

	return nil
}

func AddTaskService(newTask domain.Task) (domain.Task, error) {
	newTask.ID = primitive.NewObjectID()
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	_, err := domain.TaskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return domain.Task{}, errors.New(err.Error())
	}
	return newTask, nil
}