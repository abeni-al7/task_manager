package router

import (
	"os"

	"github.com/abeni-al7/task_manager/controllers"
	"github.com/abeni-al7/task_manager/middleware"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetTasksController)
	router.GET("/tasks/:id", controllers.GetTaskController)
	router.PUT("/tasks/:id", controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", controllers.RemoveTaskController)
	router.POST("/tasks/", controllers.AddTaskController)

	router.GET("/users", middleware.AuthMiddleware(), middleware.IsAdminMiddleware(), controllers.GetUsersController)
	router.GET("/users/:id", middleware.AuthMiddleware(), middleware.IsOwnerMiddleware(), controllers.GetUserController)
	router.PUT("/users/:id", middleware.AuthMiddleware(), middleware.IsOwnerMiddleware(), controllers.UpdateUserController)
	router.PUT("/users/:id/change-password", middleware.AuthMiddleware(), middleware.IsOwnerMiddleware(), controllers.ChangePasswordController)
	router.PUT("/promote/:id", middleware.AuthMiddleware(), middleware.IsAdminMiddleware(), controllers.PromoteUserController)
	router.DELETE("/users/:id", middleware.AuthMiddleware(), middleware.IsAdminMiddleware(), controllers.RemoveUserController)
	router.POST("/register", controllers.RegisterUserController)
	router.POST("/login", controllers.LoginUserController)

	router.Run(os.Getenv("HOST_URL"))
}