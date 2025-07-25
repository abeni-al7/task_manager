package controllers

import (
	"net/http"

	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Usecases"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase usecases.TaskUsecase
}

func (tc *TaskController) Create(ctx *gin.Context) {
	var newTask domain.Task
	
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid task")
		return
	}
	
	task, err := tc.TaskUsecase.Create(&newTask)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusCreated, task)
}

func (tc *TaskController) FetchAll(ctx *gin.Context) {
	tasks, err := tc.TaskUsecase.FetchAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) Fetch(ctx *gin.Context) {
	id := ctx.Param("id")

	task, err := tc.TaskUsecase.Fetch(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (tc *TaskController) Update(ctx *gin.Context) {
	var updatedTask domain.Task

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := tc.TaskUsecase.Update(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (tc *TaskController) Remove(ctx *gin.Context) {
	id := ctx.Param("id")

	err := tc.TaskUsecase.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}