package router

import (
	"task_manager/Delivery/controllers"
	"task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController controllers.UserController, taskController controllers.TaskController, authMiddleware Infrastructure.AuthMiddleWare) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}

	api := router.Group("/tasks")
	api.Use(authMiddleware.AuthMiddleware())
	{
		api.GET("/", taskController.GetTasks)
		api.GET("/:id", taskController.GetTaskByID)
		api.POST("/", taskController.CreateTask)
		api.PUT("/:id", taskController.UpdateTask)

		admin := api.Group("/admin")
		admin.Use(authMiddleware.AdminOnly())
		{
			admin.DELETE("/:id", taskController.DeleteTask)
		}
	}

	return router
}
