package usecases

import (
	Domain "task_manager/Domain"

	"task_manager/Repositories"
)

type TaskUsecase interface {
	CreateTask(task Domain.Task) error
	GetTaskByID(taskID string) (*Domain.Task, error)
	GetAllTasks() ([]*Domain.Task, error)
	UpdateTask(id string, task Domain.Task) error
	DeleteTask(taskID string) error
}
type taskUsecase struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepo: taskRepo,
	}
}
func (u *taskUsecase) CreateTask(task Domain.Task) error {
	return u.taskRepo.CreateTask(task)
}
func (u *taskUsecase) GetTaskByID(taskID string) (*Domain.Task, error) {
	task, err := u.taskRepo.GetTaskByID(taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func (u *taskUsecase) GetAllTasks() ([]*Domain.Task, error) {
	tasks, err := u.taskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (u *taskUsecase) UpdateTask(id string, task Domain.Task) error {
	return u.taskRepo.UpdateTask(id, task)
}
func (u *taskUsecase) DeleteTask(taskID string) error {
	return u.taskRepo.DeleteTask(taskID)
}
