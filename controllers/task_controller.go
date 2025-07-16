package controllers

import (
	"net/http"
	"strconv"

	"github.com/abeni-al7/task_manager/data"
	"github.com/abeni-al7/task_manager/models"
	"github.com/gin-gonic/gin"
)

func GetTasksController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.GetTasksService())
}

func GetTaskController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "ID must be an integer")
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "ID must be an integer")
		return
	}

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, "Not a valid Task")
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
	id, err := strconv.Atoi(idStr)
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

	task := data.AddTaskService(newTask)
	ctx.JSON(http.StatusCreated, task)
}