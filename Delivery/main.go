package main

import (
	"log"
	"os"

	"github.com/abeni-al7/task_manager/Delivery/router"
	"github.com/abeni-al7/task_manager/data"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}
	data.ConnectToMongoDB()
	routers := router.Init()
	routers.Run(os.Getenv("HOST_URL"))
}