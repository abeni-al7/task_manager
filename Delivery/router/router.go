package router

import (
	"github.com/abeni-al7/task_manager/Delivery/controllers"
	"github.com/abeni-al7/task_manager/Infrastructure"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", infrastructure.AuthMiddleware(), controllers.GetTasksController)
	router.GET("/tasks/:id", infrastructure.AuthMiddleware(), controllers.GetTaskController)
	router.PUT("/tasks/:id", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.RemoveTaskController)
	router.POST("/tasks/", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.Create())

	router.GET("/users", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.GetUsersController)
	router.GET("/users/:id", infrastructure.AuthMiddleware(), infrastructure.IsOwnerMiddleware(), controllers.GetUserController)
	router.PUT("/users/:id", infrastructure.AuthMiddleware(), infrastructure.IsOwnerMiddleware(), controllers.UpdateUserController)
	router.PUT("/users/:id/change-password", infrastructure.AuthMiddleware(), infrastructure.IsOwnerMiddleware(), controllers.ChangePasswordController)
	router.PUT("/promote/:id", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.PromoteUserController)
	router.DELETE("/users/:id", infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware(), controllers.RemoveUserController)
	router.POST("/register", controllers.RegisterUserController)
	router.POST("/login", controllers.LoginUserController)

	return router
}