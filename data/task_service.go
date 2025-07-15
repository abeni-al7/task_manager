package data

import (
	"errors"

	"github.com/abeni-al7/task_manager/models"
)

var tasks []models.Task

func GetTasks() []models.Task {
	return tasks
}

func GetTask(id int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func UpdateTask(id int, updatedTask *models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Description = updatedTask.Description
			tasks[i].Status = updatedTask.Status
			return nil
		}
	}
	return errors.New("task not found")
}

func RemoveTask(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func AddTask(newTask models.Task) {
	tasks = append(tasks, newTask)
}