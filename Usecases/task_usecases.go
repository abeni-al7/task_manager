package usecases

import (
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type TaskUsecaseInterface interface {
	Create(task *domain.Task) (domain.Task, error)
	FetchAll() ([]domain.Task, error)
	Fetch(id primitive.ObjectID) (domain.Task, error)
	Update(id primitive.ObjectID, updatedTask domain.Task) (domain.Task, error)
	Remove(id primitive.ObjectID) error
}

type TaskUsecase struct {
	taskRepo repositories.TaskRepository
}

func (tu *TaskUsecase) Create(task *domain.Task) (domain.Task, error) {
	if task.Title == "" || task.Description == "" || 
	time.Time.IsZero(task.DueDate) || task.Status == "" {
		return domain.Task{}, errors.New("missing required fields")
	}

	status := task.Status
	if status != "completed" && status != "in-progress" &&
	status != "pending" && status != "canceled" {
		return domain.Task{}, errors.New("invalid status")
	}

	task.ID = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task, err := tu.taskRepo.Create(task)
	if err != nil {
		return domain.Task{}, err
	}
	return *task, nil
}

func (tu *TaskUsecase) FetchAll() ([]domain.Task, error) {
	tasks, err := tu.taskRepo.FetchAll()
	if err != nil {
		return []domain.Task{}, err
	}
	return tasks, nil
}

func (tu *TaskUsecase) Fetch(id primitive.ObjectID) (domain.Task, error) {
	task, err := tu.taskRepo.Fetch(id)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func(tu *TaskUsecase) Update(id primitive.ObjectID, task domain.Task) (domain.Task, error) {
	status := task.Status
	if status != "completed" && status != "in-progress" &&
	status != "pending" && status != "canceled" {
		return domain.Task{}, errors.New("invalid task status value")
	}
	
	task, err := tu.taskRepo.Update(id, task)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (tu *TaskUsecase) Remove(id primitive.ObjectID) error {
	err := tu.taskRepo.Remove(id)
	return err
}