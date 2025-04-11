package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection

func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Print("Failed to connect to MongoDB")
		log.Fatal(err)
	}
	fmt.Println("Connected to Db!")

	TaskCollection = client.Database("task_management_api").Collection("tasks")
}

func GetAllTasks() []models.Task {
	findOption := options.Find()
	var tasks []models.Task
	cur, err := TaskCollection.Find(context.TODO(), bson.D{{}}, findOption)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var task models.Task
		err := cur.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task
	filter := bson.D{{Key: "id", Value: id}}
	err := TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err

	}
	return &task, nil
}
func CreateTask(task models.Task) error {
	insertResult, err := TaskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return errors.New("failed to create task")
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}
func UpdateTask(id string, updatedTask models.Task) error {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Title", Value: updatedTask.Title},
			{Key: "Description", Value: updatedTask.Description},
			{Key: "DueDate", Value: updatedTask.DueDate},
			{Key: "Status", Value: updatedTask.Status},
		}},
	}

	result, err := TaskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("failed to update task")
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")

	}
	return nil
}
func DeleteTask(id string) error {
	filter := bson.D{{Key: "id", Value: id}}
	result, err := TaskCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New("failed to delete task")
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
