package controllers

import (
	"net/http"
	domain "task_manager/Domain"
	usecases "task_manager/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController interface {
	GetTasks(c *gin.Context)
	GetTaskByID(c *gin.Context)
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}
type taskController struct {
	taskUsecase usecases.TaskUsecase
}

func NewTaskController(taskUsecase usecases.TaskUsecase) TaskController {
	return &taskController{
		taskUsecase: taskUsecase,
	}
}

func (tc *taskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server error"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *taskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *taskController) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	task.ID = uuid.New().String()
	tc.taskUsecase.CreateTask(task)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (tc *taskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := tc.taskUsecase.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (tc *taskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.taskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
