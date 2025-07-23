package main

import (
	"log"
	"os"

	"github.com/abeni-al7/task_manager/Delivery/router"
	"github.com/abeni-al7/task_manager/Domain"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}
	domain.ConnectToMongoDB()
	routers := router.Init(*domain.Database, gin.Default())
	routers.Run(os.Getenv("HOST_URL"))
}