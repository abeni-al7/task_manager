package main

import (
	"github.com/abeni-al7/task_manager/data"
	"github.com/abeni-al7/task_manager/router"
)


func main() {
	data.ConnectToMongoDB()
	router.Init()
}