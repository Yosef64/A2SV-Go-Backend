package main

import (
	"task_manager/router"
)

func main() {
	router.SetupRouter().Run(":8080")
}
