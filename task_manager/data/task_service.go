package data

import (
	"errors"
	"slices"
	"task_manager/models"
)

var tasks = []models.Task{}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}
func CreateTask(task models.Task) {
	tasks = append(tasks, task)
}
func UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}
func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = slices.Delete(tasks, i, i+1)
			return nil
		}
	}
	return errors.New("task not found")
}
