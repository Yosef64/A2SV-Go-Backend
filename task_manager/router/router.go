package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.POST("/tasks", controllers.CreateTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	return router
}
