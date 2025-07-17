package main

import (
	"log"

	"github.com/abeni-al7/task_manager/data"
	"github.com/abeni-al7/task_manager/router"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}
	data.ConnectToMongoDB()
	router.Init()
}