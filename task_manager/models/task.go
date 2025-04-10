package models

// Task represents a task in the task manager
type Task struct {
	ID          string
	Title       string
	Description string
	DueDate     string
	Status      string
}
