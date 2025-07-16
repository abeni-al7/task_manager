package data

import (
	"errors"

	"github.com/abeni-al7/task_manager/models"
)

var tasks []models.Task

func GetTasksService() []models.Task {
	return tasks
}

func GetTaskService(id int) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func UpdateTaskService(id int, updatedTask models.Task) (models.Task, error) {
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status	
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			return tasks[i], nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func RemoveTaskService(id int) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func AddTaskService(newTask models.Task) (models.Task, error) {
	id := newTask.ID
	_, err := GetTaskService(id)
	if err == nil {
		return models.Task{}, errors.New("task with this id already exists")
	}
	
	tasks = append(tasks, newTask)
	return newTask, nil
}