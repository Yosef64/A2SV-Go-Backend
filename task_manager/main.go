package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.InitMongoDB()
	router.SetupRouter().Run(":8080")
}
