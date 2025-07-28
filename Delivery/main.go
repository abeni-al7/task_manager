package main

import (
	"log"
	"os"

	"github.com/abeni-al7/task_manager/Delivery/router"
	"github.com/abeni-al7/task_manager/Repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}
	repositories.ConnectToMongoDB()
	routers := router.Init(gin.Default())
	routers.Run(os.Getenv("HOST_URL"))
}