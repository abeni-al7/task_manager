package controllers

import (
	"net/http"
	"time"

	"github.com/abeni-al7/task_manager/data"
	"github.com/abeni-al7/task_manager/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTasksController(ctx *gin.Context) {
	tasks, err := data.GetTasksService()
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func GetTaskController(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	task, err := data.GetTaskService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func UpdateTaskController(ctx *gin.Context) {
	var updatedTask models.Task

	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, "Not a valid Task")
		return
	}

	status := updatedTask.Status
	if status != "completed" && status != "in-progress" &&
	status != "pending" && status != "canceled" {
		ctx.JSON(http.StatusBadRequest, "Invalid status")
		return
	}

	task, err := data.UpdateTaskService(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func RemoveTaskController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "ID must be an integer")
		return
	}

	err = data.RemoveTaskService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func AddTaskController(ctx *gin.Context) {
	var newTask models.Task
	
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid task")
		return
	}
	
	if newTask.Title == "" || newTask.Description == "" || 
	time.Time.IsZero(newTask.DueDate) || newTask.Status == "" {
		ctx.JSON(http.StatusBadRequest, "Missing required fields")
		return
	}

	status := newTask.Status
	if status != "completed" && status != "in-progress" &&
	status != "pending" && status != "canceled" {
		ctx.JSON(http.StatusBadRequest, "Invalid status")
		return
	}
	
	task, err := data.AddTaskService(newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusCreated, task)
}