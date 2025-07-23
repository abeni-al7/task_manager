package router

import (
	"github.com/abeni-al7/task_manager/Delivery/controllers"
	"github.com/abeni-al7/task_manager/Domain"
	"github.com/abeni-al7/task_manager/Infrastructure"
	"github.com/abeni-al7/task_manager/Repositories"
	"github.com/abeni-al7/task_manager/Usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(db mongo.Database, gin *gin.Engine) *gin.Engine {
	freeRoutes := gin.Group("")
	regularRoutes := gin.Group("")
	adminRoutes := gin.Group("")
	ownerRoutes := gin.Group("")

	regularRoutes.Use(infrastructure.AuthMiddleware())
	adminRoutes.Use(infrastructure.AuthMiddleware(), infrastructure.IsAdminMiddleware())
	ownerRoutes.Use(infrastructure.AuthMiddleware(), infrastructure.IsOwnerMiddleware())

	AuthRouter(db, freeRoutes)
	TaskAccessRouter(db, regularRoutes)
	TaskManipulationRouter(db, adminRoutes)
	UserControlRouter(db, adminRoutes)
	AccountControlRouter(db, ownerRoutes)
	return gin
}

func AuthRouter(db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(*ur),
	}

	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
}

func TaskAccessRouter(db mongo.Database, group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(db, domain.TaskCollection)
	tc := &controllers.TaskController{
		TaskUsecase: *usecases.NewTaskUsecase(*tr),
	}

	group.GET("/tasks", tc.FetchAll)
	group.GET("/tasks/:id", tc.Fetch)
}

func TaskManipulationRouter(db mongo.Database, group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(db, domain.TaskCollection)
	tc := &controllers.TaskController{
		TaskUsecase: *usecases.NewTaskUsecase(*tr),
	}

	group.PUT("/tasks/:id", tc.Update)
	group.DELETE("/tasks/:id", tc.Remove)
	group.POST("/tasks", tc.Create)
}

func UserControlRouter(db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(*ur),
	}

	group.GET("/users", uc.FetchAll)
	group.PUT("/promote/:id", uc.Promote)
	group.DELETE("/users/:id", uc.Remove)
}

func AccountControlRouter(db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, domain.UserCollection)
	uc := &controllers.UserController{
		UserUsecase: *usecases.NewUserUsecase(*ur),
	}

	group.GET("/users/:id", uc.Fetch)
	group.PUT("/users/:id", uc.Update)
	group.PUT("/users/:id/change-password", uc.ChangePassword)
}