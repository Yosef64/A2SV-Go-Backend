package Repositories

import (
	"context"
	"errors"
	Domain "task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(task Domain.Task) error
	GetAllTasks() ([]*Domain.Task, error)
	GetTaskByID(userID string) (*Domain.Task, error)
	UpdateTask(id string, task Domain.Task) error
	DeleteTask(taskID string) error
}

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &taskRepository{collection: collection}
}

func (r *taskRepository) CreateTask(task Domain.Task) error {
	_, err := r.collection.InsertOne(context.TODO(), task)
	return err
}

func (r *taskRepository) GetAllTasks() ([]*Domain.Task, error) {
	cur, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var tasks []*Domain.Task
	for cur.Next(context.TODO()) {
		var task Domain.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *taskRepository) GetTaskByID(id string) (*Domain.Task, error) {
	var task Domain.Task
	filter := bson.D{{Key: "id", Value: id}}
	err := r.collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err

	}
	return &task, nil
}

func (r *taskRepository) UpdateTask(id string, task Domain.Task) error {
	task.ID = id
	filter := bson.M{"id": id}
	update := bson.D{{Key: "$set", Value: task}}

	result, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (r *taskRepository) DeleteTask(taskID string) error {
	filter := bson.M{"id": taskID}
	result, err := r.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found or unauthorized")
	}
	return nil
}
