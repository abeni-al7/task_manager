package controllers

import (
	"net/http"
	"strconv"

	"github.com/abeni-al7/task_manager/data"
	"github.com/gin-gonic/gin"
)

func GetTasksController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data.GetTasksService())
}

func GetTaskController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	task, err := data.GetTaskService(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, task)
}