package router

import (
	"github.com/abeni-al7/task_manager/Delivery/controllers"
	domain "github.com/abeni-al7/task_manager/Domain"
	infrastructure "github.com/abeni-al7/task_manager/Infrastructure"
	repositories "github.com/abeni-al7/task_manager/Repositories"
	usecases "github.com/abeni-al7/task_manager/Usecases"
	"github.com/gin-gonic/gin"
)

func Init(gin *gin.Engine) *gin.Engine {
	freeRoutes := gin.Group("")
	regularRoutes := gin.Group("")
	adminRoutes := gin.Group("")
	ownerRoutes := gin.Group("")

	regularRoutes.Use(infrastructure.AuthMiddleware())
	adminRoutes.Use(infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware())
	ownerRoutes.Use(infrastructure.AuthMiddleware(), infrastructure.IsOwnerMiddleware())

	AuthRouter(freeRoutes)
	TaskAccessRouter(regularRoutes)
	TaskManipulationRouter(adminRoutes)
	UserControlRouter(adminRoutes)
	AccountControlRouter(ownerRoutes)
	return gin
}

func AuthRouter(group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(ur, new(infrastructure.Infrastructure)),
	}

	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
}

func TaskAccessRouter(group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(domain.TaskCollection)
	tc := &controllers.TaskController{
		TaskUsecase: *usecases.NewTaskUsecase(tr),
	}

	group.GET("/tasks", tc.FetchAll)
	group.GET("/tasks/:id", tc.Fetch)
}

func TaskManipulationRouter(group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(domain.TaskCollection)
	tc := &controllers.TaskController{
		TaskUsecase: *usecases.NewTaskUsecase(tr),
	}

	group.PUT("/tasks/:id", tc.Update)
	group.DELETE("/tasks/:id", tc.Remove)
	group.POST("/tasks", tc.Create)
}

func UserControlRouter(group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(ur, new(infrastructure.Infrastructure)),
	}

	group.GET("/users", uc.FetchAll)
	group.PUT("/promote/:id", uc.Promote)
	group.DELETE("/users/:id", uc.Remove)
}

func AccountControlRouter(group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(ur, new(infrastructure.Infrastructure)),
	}

	group.GET("/users/:id", uc.Fetch)
	group.PUT("/users/:id", uc.Update)
	group.PUT("/users/:id/change-password", uc.ChangePassword)
}