package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type TaskRepositoryInterface interface {
	Create(task *domain.Task) (domain.Task, error)
	FetchAll() ([]domain.Task, error)
	Fetch(id primitive.ObjectID) (domain.Task, error)
	Update(id primitive.ObjectID, task domain.Task) (domain.Task, error)
	Remove(id primitive.ObjectID) error
}

type TaskRepository struct {
	database mongo.Database
	collection string
}

func (tr *TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
	_, err :=tr.database.Collection(tr.collection).InsertOne(context.TODO(), task)
	if err != nil {
		return &domain.Task{}, errors.New("cannot insert task to database")
	}
	return task, nil
}

func (tr *TaskRepository) FetchAll() ([]domain.Task, error) {
	var tasks []domain.Task

	cur, err := tr.database.Collection(tr.collection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []domain.Task{}, errors.New("cannot retrieve tasks")
	}

	err = cur.All(context.TODO(), tasks)
	if err != nil {
		return []domain.Task{}, errors.New("cannot retrieve tasks")
	}

	cur.Close(context.TODO())

	return tasks, nil
}

func(tr *TaskRepository) Fetch(id primitive.ObjectID) (domain.Task, error) {
	var task domain.Task

	filter := bson.D{{Key: "_id", Value: id}}

	err := tr.database.Collection(tr.collection).FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}

	return task, nil
}

func(tr *TaskRepository) Update(id primitive.ObjectID, task domain.Task) (domain.Task, error) {
	var updatedTask domain.Task
	filter := bson.D{{Key: "_id", Value: id}}

	fields := bson.D{}
	if task.Title != "" {
		fields = append(fields, bson.E{Key: "title", Value: task.Title})
	}
	if task.Description != "" {
		fields = append(fields, bson.E{Key: "description", Value: task.Description})
	}
	if !time.Time.IsZero(task.DueDate) {
		fields = append(fields, bson.E{Key: "due_date", Value: task.DueDate})
	}
	if task.Status != "" {
		fields = append(fields, bson.E{Key: "status", Value: task.Status})
	}
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: fields}}

	_, err := tr.database.Collection(tr.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return domain.Task{}, errors.New(err.Error())
	}
	
	err = tr.database.Collection(tr.collection).FindOne(context.TODO(), filter).Decode(&updatedTask)
	if err != nil {
		return domain.Task{}, errors.New("task not found")
	}
	return updatedTask, nil
}

func (tr *TaskRepository) Remove(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id", Value: id}}
	_, err := tr.database.Collection(tr.collection).DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("task not found")
	}

	return nil
}