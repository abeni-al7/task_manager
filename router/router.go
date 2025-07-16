package router

import (
	"github.com/abeni-al7/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasksController)
	router.GET("/tasks/:id", controllers.GetTaskController)
	router.PUT("/tasks/:id", controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", controllers.RemoveTaskController)
	router.POST("/tasks/:id", controllers.AddTaskController)

	router.Run("localhost:8080")
}