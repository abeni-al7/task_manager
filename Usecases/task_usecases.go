package usecases

import (
	"errors"
	"time"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Usecases/interfaces"
)

type TaskUsecase struct {
	taskRepo usecases.ITaskRepo
}

func NewTaskUsecase(tr usecases.ITaskRepo) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: tr,
	}
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

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	newTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return domain.Task{}, err
	}
	return newTask, nil
}

func (tu *TaskUsecase) FetchAll() ([]domain.Task, error) {
	tasks, err := tu.taskRepo.FetchAll()
	if err != nil {
		return []domain.Task{}, err
	}
	return tasks, nil
}

func (tu *TaskUsecase) Fetch(id string) (domain.Task, error) {
	task, err := tu.taskRepo.Fetch(id)
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func(tu *TaskUsecase) Update(id string, task domain.Task) (domain.Task, error) {
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

func (tu *TaskUsecase) Remove(id string) error {
	err := tu.taskRepo.Remove(id)
	return err
}